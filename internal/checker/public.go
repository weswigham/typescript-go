package checker

import (
	"github.com/microsoft/typescript-go/internal/ast"
	"github.com/microsoft/typescript-go/internal/core"
	"github.com/microsoft/typescript-go/internal/diagnostics"
)

func (c *Checker) GetMissingType() *Type {
	return c.missingType
}

func (c *Checker) GetUnknownType() *Type {
	return c.unknownType
}

func (c *Checker) GetErrorType() *Type {
	return c.errorType
}

func (c *Checker) GetAnyType() *Type {
	return c.anyType
}

func (c *Checker) GetGlobalArrayType() *Type {
	return c.globalArrayType
}

func (c *Checker) GetGlobalReadonlyArrayType() *Type {
	return c.globalReadonlyArrayType
}

func (c *Checker) GetIntrinsicMarkerType() *Type {
	return c.intrinsicMarkerType
}

func (c *Checker) GetMarkerSuperTypeForCheck() *Type {
	return c.markerSuperTypeForCheck
}

func (c *Checker) GetMarkerSubTypeForCheck() *Type {
	return c.markerSubTypeForCheck
}

func (c *Checker) GetVarianceTypeParameter() *Type {
	return c.varianceTypeParameter
}

func (c *Checker) GetUnknownSymbol() *ast.Symbol {
	return c.unknownSymbol
}

func (c *Checker) GetOptionalType(t *Type, isProperty bool) *Type {
	return c.getOptionalType(t, isProperty)
}

func (c *Checker) GetTypeWithFacts(t *Type, include TypeFacts) *Type {
	return c.getTypeWithFacts(t, include)
}

func (c *Checker) GetTypeFromTypeNode(node *ast.Node) *Type {
	return c.getTypeFromTypeNode(node)
}

func (c *Checker) TryGetResolvedSymbolOfTypeNode(existing *ast.Node) *ast.Symbol {
	// `type` is a reference type, and `existing` is a type reference node, but we still need to make sure they refer to the _same_ target type
	// before we go comparing their type argument counts.
	c.getTypeFromTypeNode(existing)
	// call to ensure symbol is resolved
	links := c.symbolNodeLinks.TryGet(existing)
	if links == nil {
		return nil
	}
	return links.resolvedSymbol
}

func (c *Checker) GetDeclaredTypeOfSymbol(symbol *ast.Symbol) *Type {
	return c.getDeclaredTypeOfSymbol(symbol)
}

func (c *Checker) GetResolvedTypeWithoutAbstractConstructSignatures(type_ *Type) *Type {
	if type_.Flags()&TypeFlagsStructuredType == 0 {
		return type_
	}
	t := c.resolveStructuredTypeMembers(type_)
	if len(t.ConstructSignatures()) == 0 {
		return t.AsType()
	}
	if t.objectTypeWithoutAbstractConstructSignatures != nil {
		return t.objectTypeWithoutAbstractConstructSignatures
	}
	constructSignatures := core.Filter(t.ConstructSignatures(), func(signature *Signature) bool {
		return signature.flags&SignatureFlagsAbstract == 0
	})
	if len(constructSignatures) == len(t.ConstructSignatures()) {
		t.objectTypeWithoutAbstractConstructSignatures = t.AsType()
		return t.AsType()
	}
	typeCopy := c.newAnonymousType(t.symbol, t.members, t.CallSignatures(), core.IfElse(len(constructSignatures) > 0, constructSignatures, []*Signature{}), t.indexInfos)
	t.objectTypeWithoutAbstractConstructSignatures = typeCopy
	typeCopy.AsStructuredType().objectTypeWithoutAbstractConstructSignatures = typeCopy
	return typeCopy
}

func (c *Checker) TryGetNameTypeOfSymbol(symbol *ast.Symbol) *Type {
	if c.valueSymbolLinks.Has(symbol) {
		return c.valueSymbolLinks.Get(symbol).nameType
	}
	return nil
}

func (c *Checker) GetCompilerOptions() *core.CompilerOptions {
	return c.compilerOptions
}

func (c *Checker) GetEmitModuleFormatOfFile(sourceFile *ast.SourceFile) core.ModuleKind {
	return c.program.GetEmitModuleFormatOfFile(sourceFile)
}

func (c *Checker) GetExportsOfSymbol(symbol *ast.Symbol) ast.SymbolTable {
	return c.getExportsOfSymbol(symbol)
}

func (c *Checker) GetSymbolIfSameReference(s1 *ast.Symbol, s2 *ast.Symbol) *ast.Symbol {
	return c.getSymbolIfSameReference(s1, s2)
}

func (c *Checker) SortSymbols(symbols []*ast.Symbol) {
	c.sortSymbols(symbols)
}

func (c *Checker) CompareSymbols(s1 *ast.Symbol, s2 *ast.Symbol) int {
	return c.compareSymbols(s1, s2)
}

func (c *Checker) GetMembersOfSymbol(symbol *ast.Symbol) ast.SymbolTable {
	return c.getMembersOfSymbol(symbol)
}

// The full set of type parameters for a generic class or interface type consists of its outer type parameters plus
// its locally declared type parameters.
func (c *Checker) GetTypeParametersOfClassOrInterface(symbol *ast.Symbol) []*Type {
	result := make([]*Type, 0)
	result = append(result, c.getOuterTypeParametersOfClassOrInterface(symbol)...)
	result = append(result, c.getLocalTypeParametersOfClassOrInterfaceOrTypeAlias(symbol)...)
	return result
}

func (c *Checker) TryGetMapperOfSymbol(symbol *ast.Symbol) *TypeMapper {
	if c.valueSymbolLinks.Has(symbol) {
		return c.valueSymbolLinks.Get(symbol).mapper
	}
	return nil
}

func (c *Checker) GetSymbolChain(
	symbol *ast.Symbol,
	meaning ast.SymbolFlags,
	endOfChain bool,
	yieldModuleSymbol bool,
	enclosingDeclaration *ast.Node,
	useOnlyExternalAliasing bool,
	compareSymbols func(s1 *ast.Symbol, s2 *ast.Symbol) int,
) []*ast.Symbol {
	return c.getSymbolChain(
		symbol,
		meaning,
		endOfChain,
		yieldModuleSymbol,
		enclosingDeclaration,
		useOnlyExternalAliasing,
		compareSymbols,
	)
}

func (ch *Checker) GetFileSymbolIfFileSymbolExportEqualsContainer(d *ast.Node, container *ast.Symbol) *ast.Symbol {
	return ch.getFileSymbolIfFileSymbolExportEqualsContainer(d, container)
}

func (c *Checker) GetDefaultFromTypeParameter(t *Type) *Type {
	return c.getDefaultFromTypeParameter(t)
}

func (c *Checker) ResolveName(location *ast.Node, name string, meaning ast.SymbolFlags, nameNotFoundMessage *diagnostics.Message, isUse bool, excludeGlobals bool) *ast.Symbol {
	return c.resolveName(location, name, meaning, nameNotFoundMessage, isUse, excludeGlobals)
}

func (c *Checker) GetHomomorphicTypeVariable(t *Type) *Type {
	return c.getHomomorphicTypeVariable(t)
}

func (c *Checker) IsMappedTypeWithKeyofConstraintDeclaration(t *Type) bool {
	return c.isMappedTypeWithKeyofConstraintDeclaration(t)
}

func (c *Checker) GetModifiersTypeFromMappedType(t *Type) *Type {
	return c.getModifiersTypeFromMappedType(t)
}

func (c *Checker) GetConstraintTypeFromMappedType(t *Type) *Type {
	return c.getConstraintTypeFromMappedType(t)
}

func (c *Checker) GetTypeParameterFromMappedType(t *Type) *Type {
	return c.getTypeParameterFromMappedType(t)
}

func (c *Checker) GetTemplateTypeFromMappedType(t *Type) *Type {
	return c.getTemplateTypeFromMappedType(t)
}

func (c *Checker) GetNameTypeFromMappedType(t *Type) *Type {
	return c.getNameTypeFromMappedType(t)
}

func (c *Checker) GetConstraintOfTypeParameter(typeParameter *Type) *Type {
	return c.getConstraintOfTypeParameter(typeParameter)
}

func (c *Checker) RemoveMissingType(t *Type, isOptional bool) *Type {
	return c.removeMissingType(t, isOptional)
}

func (c *Checker) NewTypeParameter() *Type {
	return c.newTypeParameter(
		c.newSymbol(ast.SymbolFlagsTypeParameter, "T"),
	)
}

func (c *Checker) InstantiateType(t *Type, m *TypeMapper) *Type {
	return c.instantiateType(t, m)
}

func (c *Checker) GetTargetSymbol(s *ast.Symbol) *ast.Symbol {
	return c.getTargetSymbol(s)
}

func (c *Checker) GetLocalTypeParametersOfClassOrInterfaceOrTypeAlias(symbol *ast.Symbol) []*Type {
	return c.getLocalTypeParametersOfClassOrInterfaceOrTypeAlias(symbol)
}

func (c *Checker) GetConstraintDeclaration(t *Type) *ast.Node {
	return c.getConstraintDeclaration(t)
}

func (c *Checker) GetTypeOfSymbol(symbol *ast.Symbol) *Type {
	return c.getTypeOfSymbol(symbol)
}

func (c *Checker) IsOptionalParameter(node *ast.Node) bool {
	return c.isOptionalParameter(node)
}

func (c *Checker) isOptionalParameter(node *ast.Node) bool {
	// !!! TODO: JSDoc support
	// if (hasEffectiveQuestionToken(node)) {
	// 	return true;
	// }
	if ast.IsParameter(node) && node.AsParameterDeclaration().QuestionToken != nil {
		return true
	}
	if !ast.IsParameter(node) {
		return false
	}
	if node.Initializer() != nil {
		signature := c.getSignatureFromDeclaration(node.Parent)
		parameterIndex := core.FindIndex(node.Parent.Parameters(), func(p *ast.ParameterDeclarationNode) bool { return p == node })
		// Debug.assert(parameterIndex >= 0); // !!!
		// Only consider syntactic or instantiated parameters as optional, not `void` parameters as this function is used
		// in grammar checks and checking for `void` too early results in parameter types widening too early
		// and causes some noImplicitAny errors to be lost.
		return parameterIndex >= c.getMinArgumentCountEx(signature, MinArgumentCountFlagsStrongArityForUntypedJS|MinArgumentCountFlagsVoidIsNonOptional)
	}
	iife := ast.GetImmediatelyInvokedFunctionExpression(node.Parent)
	if iife != nil {
		parameterIndex := core.FindIndex(node.Parent.Parameters(), func(p *ast.ParameterDeclarationNode) bool { return p == node })
		return node.Type() == nil &&
			node.AsParameterDeclaration().DotDotDotToken == nil &&
			parameterIndex >= len(c.getEffectiveCallArguments(iife))
	}

	return false
}

func (c *Checker) InstantiateTypePredicate(predicate *TypePredicate, mapper *TypeMapper) *TypePredicate {
	return c.instantiateTypePredicate(predicate, mapper)
}

func (c *Checker) GetTypePredicateOfSignature(sig *Signature) *TypePredicate {
	return c.getTypePredicateOfSignature(sig)
}

func (c *Checker) GetWidenedType(t *Type) *Type {
	return c.getWidenedType(t)
}

func (c *Checker) GetRegularTypeOfExpression(expr *ast.Node) *Type {
	return c.getRegularTypeOfExpression(expr)
}

func (c *Checker) GetExpandedParametersEx(sig *Signature, skipUnionExpanding bool) [][]*ast.Symbol {
	return c.getExpandedParameters(sig, skipUnionExpanding)
}

func (c *Checker) GetReturnTypeOfSignature(sig *Signature) *Type {
	return c.getReturnTypeOfSignature(sig)
}

func (c *Checker) GetSymbolOfDeclaration(node *ast.Node) *ast.Symbol {
	return c.getSymbolOfDeclaration(node)
}

func (c *Checker) GetWriteTypeOfSymbol(symbol *ast.Symbol) *Type {
	return c.getWriteTypeOfSymbol(symbol)
}

func (c *Checker) GetWidenedLiteralType(t *Type) *Type {
	return c.getWidenedLiteralType(t)
}

func (c *Checker) TryGetReverseMappedSymbolPropertyType(symbol *ast.Symbol) *Type {
	if c.ReverseMappedSymbolLinks.Has(symbol) {
		links := c.ReverseMappedSymbolLinks.TryGet(symbol)
		return links.propertyType
	}
	return nil
}

func (c *Checker) TryGetReverseMappedSymbolMappedType(symbol *ast.Symbol) *Type {
	if c.ReverseMappedSymbolLinks.Has(symbol) {
		links := c.ReverseMappedSymbolLinks.TryGet(symbol)
		return links.mappedType
	}
	return nil
}

func (c *Checker) RequiresAddingImplicitUndefined(declaration *ast.Node, symbol *ast.Symbol, enclosingDeclaration *ast.Node) bool {
	if !ast.IsParseTreeNode(declaration) {
		return false
	}
	switch declaration.Kind {
	case ast.KindPropertyDeclaration, ast.KindPropertySignature, ast.KindJSDocPropertyTag:
		if symbol == nil {
			symbol = c.getSymbolOfDeclaration(declaration)
		}
		type_ := c.getTypeOfSymbol(symbol)
		return !!((symbol.Flags&ast.SymbolFlagsProperty != 0) && (symbol.Flags&ast.SymbolFlagsOptional != 0) && isOptionalDeclaration(declaration) && c.ReverseMappedSymbolLinks.Has(symbol) && c.ReverseMappedSymbolLinks.Get(symbol).mappedType != nil && containsNonMissingUndefinedType(c, type_))
	case ast.KindParameter, ast.KindJSDocParameterTag:
		return c.requiresAddingImplicitUndefined(declaration, enclosingDeclaration)
	default:
		panic("Node cannot possibly require adding undefined")
	}
}

func (c *Checker) requiresAddingImplicitUndefined(parameter *ast.Node, enclosingDeclaration *ast.Node) bool {
	return (c.isRequiredInitializedParameter(parameter, enclosingDeclaration) || c.isOptionalUninitializedParameterProperty(parameter)) && !c.declaredParameterTypeContainsUndefined(parameter)
}

func (c *Checker) declaredParameterTypeContainsUndefined(parameter *ast.Node) bool {
	// typeNode := getNonlocalEffectiveTypeAnnotationNode(parameter); // !!! JSDoc Support
	typeNode := parameter.Type()
	if typeNode == nil {
		return false
	}
	type_ := c.getTypeFromTypeNode(typeNode)
	// allow error type here to avoid confusing errors that the annotation has to contain undefined when it does in cases like this:
	//
	// export function fn(x?: Unresolved | undefined): void {}
	return c.isErrorType(type_) || c.containsUndefinedType(type_)
}

func (c *Checker) isOptionalUninitializedParameterProperty(parameter *ast.Node) bool {
	return c.strictNullChecks &&
		c.IsOptionalParameter(parameter) &&
		( /*isJSDocParameterTag(parameter) ||*/ parameter.Initializer() != nil) && // !!! TODO: JSDoc support
		ast.HasSyntacticModifier(parameter, ast.ModifierFlagsParameterPropertyModifier)
}

func (c *Checker) isRequiredInitializedParameter(parameter *ast.Node, enclosingDeclaration *ast.Node) bool {
	if c.strictNullChecks || c.IsOptionalParameter(parameter) || /*isJSDocParameterTag(parameter) ||*/ parameter.Initializer() == nil { // !!! TODO: JSDoc Support
		return false
	}
	if ast.HasSyntacticModifier(parameter, ast.ModifierFlagsParameterPropertyModifier) {
		return enclosingDeclaration != nil && ast.IsFunctionLikeDeclaration(enclosingDeclaration)
	}
	return true
}

func (c *Checker) GetTypeOfExpression(node *ast.Node) *Type {
	return c.getTypeOfExpression(node)
}

func (c *Checker) GetNonMissingTypeOfSymbol(symbol *ast.Symbol) *Type {
	return c.getNonMissingTypeOfSymbol(symbol)
}

func (c *Checker) HasLateBindableName(node *ast.Node) bool {
	return c.hasLateBindableName(node)
}

func (c *Checker) IsErrorType(t *Type) bool {
	return c.isErrorType(t)
}

func (c *Checker) GetSignatureFromDeclaration(declaration *ast.Node) *Signature {
	return c.getSignatureFromDeclaration(declaration)
}

func (c *Checker) GetSignaturesOfType(t *Type, kind SignatureKind) []*Signature {
	return c.getSignaturesOfType(t, kind)
}

func (c *Checker) GetPropertiesOfObjectType(t *Type) []*ast.Symbol {
	return c.getPropertiesOfObjectType(t)
}

func (c *Checker) FilterType(t *Type, f func(*Type) bool) *Type {
	return c.filterType(t, f)
}

func (c *Checker) IsReadonlySymbol(symbol *ast.Symbol) bool {
	return c.isReadonlySymbol(symbol)
}

func (c *Checker) IsGenericMappedType(t *Type) bool {
	return c.isGenericMappedType(t)
}

func (c *Checker) ResolveStructuredTypeMembers(t *Type) *StructuredType {
	return c.resolveStructuredTypeMembers(t)
}

func (c *Checker) GetOrCreateTypeFromSignature(sig *Signature, outerTypeParameters []*Type) *Type {
	return c.getOrCreateTypeFromSignature(sig, outerTypeParameters)
}

func (c *Checker) GetIntersectionType(types []*Type) *Type {
	return c.getIntersectionType(types)
}

func (c *Checker) IsClassInstanceSide(t *Type) bool {
	return t.symbol != nil && t.symbol.Flags&ast.SymbolFlagsClass != 0 && (t == c.getDeclaredTypeOfClassOrInterface(t.symbol) || (t.flags&TypeFlagsObject != 0 && t.objectFlags&ObjectFlagsIsClassInstanceClone != 0))
}

func (c *Checker) GetBaseTypeVariableOfClass(symbol *ast.Symbol) *Type {
	return c.getBaseTypeVariableOfClass(symbol)
}

func (c *Checker) GetTrueTypeFromConditionalType(t *Type) *Type {
	return c.getTrueTypeFromConditionalType(t)
}

func (c *Checker) GetFalseTypeFromConditionalType(t *Type) *Type {
	return c.getFalseTypeFromConditionalType(t)
}

func (c *Checker) GetSymbolOfNode(node *ast.Node) *ast.Symbol {
	return c.getSymbolOfNode(node)
}

func (c *Checker) GetTypeArguments(t *Type) []*Type {
	return c.getTypeArguments(t)
}

func (c *Checker) GetTypeReferenceArity(t *Type) int {
	return c.getTypeReferenceArity(t)
}

func (c *Checker) GetTupleElementLabel(elementInfo TupleElementInfo, restSymbol *ast.Symbol, index int) string {
	return c.getTupleElementLabel(elementInfo, restSymbol, index)
}

func (c *Checker) GetGlobalIterableType() *Type {
	return c.getGlobalIterableType()
}

func (c *Checker) GetGlobalIterableIteratorType() *Type {
	return c.getGlobalIterableIteratorType()
}

func (c *Checker) GetGlobalAsyncIterableType() *Type {
	return c.getGlobalAsyncIterableType()
}

func (c *Checker) GetGlobalAsyncIterableIteratorType() *Type {
	return c.getGlobalAsyncIterableIteratorType()
}

func (c *Checker) IsReferenceToType(t *Type, target *Type) bool {
	return c.isReferenceToType(t, target)
}

func (c *Checker) IsTypeIdenticalTo(source *Type, target *Type) bool {
	return c.isTypeIdenticalTo(source, target)
}

func (c *Checker) GetReducedType(t *Type) *Type {
	return c.getReducedType(t)
}

func (c *Checker) GetParentOfSymbol(symbol *ast.Symbol) *ast.Symbol {
	return c.getParentOfSymbol(symbol)
}

func (c *Checker) GetInferredTypeParameterConstraint(t *Type, omitTypeReferences bool) *Type {
	return c.getInferredTypeParameterConstraint(t, omitTypeReferences)
}

func (c *Checker) FormatUnionTypes(types []*Type) []*Type {
	return c.formatUnionTypes(types)
}

func (c *Checker) IsNoInferType(t *Type) bool {
	return c.isNoInferType(t)
}

func (c *Checker) GetGlobalTypeAliasSymbol(name string, arity int, reportErrors bool) *ast.Symbol {
	return c.getGlobalTypeAliasSymbol(name, arity, reportErrors)
}
