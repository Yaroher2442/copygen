package parser

import (
	"fmt"
	"go/types"

	"github.com/fatih/structtag"
	"github.com/switchupcb/copygen/cli/models"
	"github.com/switchupcb/copygen/cli/parser/options"
)

// fieldParser represents the parameters required to parse a types.Type into a *models.Field.
type fieldParser struct {
	// field represents the current field being built.
	field *models.Field

	// parent represents the parent of the field parse.
	parent *models.Field

	// cyclic is a key value cache used to prevent cyclic fields from unnecessary duplication or stack overflow.
	cyclic map[string]*models.Field

	// options represents the field options defined above the models.Function
	options []*options.Option
}

// parseField parses a types.Type into a *models.Field recursively.
func (fp fieldParser) parseField(typ types.Type) *models.Field {
	if fp.field == nil {
		fp.field = &models.Field{Parent: fp.parent}
	}

	switch x := typ.(type) {

	// Named Types (Alias)
	// https://go.googlesource.com/example/+/HEAD/gotypes#named-types
	case *types.Named:
		setFieldImportAndPackage(fp.field, x.Obj().Pkg())
		setDefinition(fp.field, x.Obj().Name())

		// A named type is either:
		//   1. an alias (i.e `Placeholder` in `type Placeholder bool`)
		//   2. a struct (i.e `Account` in `type Account struct`)
		//   3. an interface (i.e `error` in `type error interface`)
		//   4. a collected type (i.e `domain.Account` in `[]domain.Account`)
		//
		// Underlying types are only important in case 3,
		// when we need to parse extra information from the field.
		if !fp.field.IsCollection() {

			// *types.Basic must handle case 1.
			return fp.parseField(x.Underlying())
		}

	// Basic Types
	// https://go.googlesource.com/example/+/HEAD/gotypes#basic-types
	case *types.Basic:
		if !fp.field.IsAlias() {
			fp.field.Package = ""
			setDefinition(fp.field, x.Name())
		}

	// Simple Composite Types
	// https://go.googlesource.com/example/+/HEAD/gotypes#simple-composite-types
	case *types.Pointer:
		if fp.field.Definition == "" && fp.field.Pointer == "" {
			fp.field.Pointer = models.Pointer
		} else {
			setDefinition(fp.field, models.CollectionPointer)
		}
		return fp.parseField(x.Elem())

	case *types.Array:
		setDefinition(fp.field, "["+fmt.Sprint(x.Len())+"]")
		return fp.parseField(x.Elem())

	case *types.Slice:
		setDefinition(fp.field, models.CollectionSlice)
		return fp.parseField(x.Elem())

	case *types.Map:
		setDefinition(fp.field, models.CollectionMap+"[")
		_ = fp.parseField(x.Key())
		setDefinition(fp.field, "]")
		return fp.parseField(x.Elem())

	case *types.Chan:
		setDefinition(fp.field, models.CollectionChan+" ")
		return fp.parseField(x.Elem())

	// Function (without Receivers)
	// https://go.googlesource.com/example/+/HEAD/gotypes#function-and-method-types
	case *types.Signature:

		// set the parameters.
		setDefinition(fp.field, models.CollectionFunc+"(")
		for i := 0; i < x.Params().Len(); i++ {
			_ = fp.parseField(x.Params().At(i).Type())
			if i+1 != x.Params().Len() {
				setDefinition(fp.field, ", ")
			}
		}
		setDefinition(fp.field, ") ")

		// set the results.
		if x.Results().Len() > 1 {
			setDefinition(fp.field, "(")
		}
		for i := 0; i < x.Results().Len(); i++ {
			_ = fp.parseField(x.Results().At(i).Type())
			if i+1 != x.Results().Len() {
				setDefinition(fp.field, ", ")
			}
		}
		if x.Results().Len() > 1 {
			setDefinition(fp.field, ")")
		}

	// Interface Types
	// https://go.googlesource.com/example/+/HEAD/gotypes#interface-types
	case *types.Interface:
		if !fp.field.IsAlias() {
			if x.Empty() {
				setDefinition(fp.field, x.String())
			} else {
				setDefinition(fp.field, models.CollectionInterface+"{")
				for i := 0; i < x.NumMethods(); i++ {
					_ = fp.parseField(x.Method(i).Type())
					setDefinition(fp.field, "; ")
				}
				setDefinition(fp.field, "}")
			}
		}

	// Struct Types
	// https://go.googlesource.com/example/+/HEAD/gotypes#struct-types
	case *types.Struct:
		for i := 0; i < x.NumFields(); i++ {
			if subfield, ok := fp.cyclic[x.Field(i).String()]; ok {
				fp.field.Fields = append(fp.field.Fields, subfield)
				continue
			}

			subfield := fp.parseSubfield(x.Field(i), x.Tag(i))
			fp.field.Fields = append(fp.field.Fields, subfield)
		}
	}

	// set the field's options.
	options.SetFieldOptions(fp.field, fp.options)
	filterFieldDepth(fp.field, fp.field.Options.Depth, 0)

	setFieldVariableName(fp.field, "."+alphastring(fp.field.Definition))
	fp.cyclic[typ.String()] = fp.field
	return fp.field
}

// parseSubfield parses a types.Var into a *models.Field.
func (fp fieldParser) parseSubfield(x *types.Var, tag string) *models.Field {
	subfield := &models.Field{
		VariableName: "." + x.Name(),
		Name:         x.Name(),
		Parent:       fp.field,
	}
	setFieldImportAndPackage(subfield, x.Pkg())
	setTags(subfield, tag)
	subfieldParser := &fieldParser{
		field:   subfield,
		parent:  nil,
		options: fp.options,
		cyclic:  fp.cyclic,
	}

	fp.cyclic[x.String()] = subfield
	subfield = subfieldParser.parseField(x.Type())
	return subfield
}

// setFieldImportAndPackage sets the import and package of a field.
func setFieldImportAndPackage(field *models.Field, pkg *types.Package) {
	if pkg == nil {
		return
	}

	field.Import = pkg.Path()
	if ignorepkgpath != field.Import {
		if _, ok := aliasImportMap[field.Import]; ok {
			field.Package = aliasImportMap[field.Import]
		} else {
			field.Package = pkg.Name()
		}
	}

	// field collections set collected types' packages in the field.Definition.
	// i.e map[*domain.Account]string
	if field.IsCollection() {
		setDefinition(field, field.Package+".")
		field.Package = ""
	}
}

// setFieldVariableName sets a field's variable name.
func setFieldVariableName(field *models.Field, varname string) {
	if field.VariableName == "" {
		field.VariableName = varname
	}
}

// setDefinition sets a field's definition.
func setDefinition(field *models.Field, def string) {
	switch {
	case field.Definition == "":
		field.Definition = def
	case field.IsInterface():
		field.Definition += def
	case field.IsCollection():
		field.Definition += def
	}
}

// setTags sets the tags for a field.
func setTags(field *models.Field, rawtag string) {
	// rawtag represents tags as they are defined (i.e `api:"id", json:"tag"`).
	tags, err := structtag.Parse(rawtag)
	if err != nil {
		fmt.Printf("WARNING: could not parse tag for field %v\n%v", field.FullName(""), err)
	}

	if field.Tags == nil {
		field.Tags = make(map[string]map[string][]string, tags.Len())
	}

	for _, tag := range tags.Tags() {
		field.Tags[tag.Key] = map[string][]string{
			tag.Name: tag.Options,
		}
	}
}

// filterFieldDepth filters a field's fields according to it's depth level.
func filterFieldDepth(field *models.Field, maxdepth, curdepth int) {
	if maxdepth == 0 {
		return
	}

	if maxdepth < 0 || maxdepth <= curdepth {
		field.Fields = make([]*models.Field, 0)
		return
	}

	for _, f := range field.Fields {
		filterFieldDepth(f, maxdepth+f.Options.Depth, curdepth+1)
	}
}

// alphastring only returns alphabetic characters (English) in a string.
func alphastring(s string) string {
	bytes := []byte(s)
	i := 0
	for _, b := range bytes {
		if ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || b == ' ' {
			bytes[i] = b
			i++
		}
	}

	return string(bytes[:i])
}
