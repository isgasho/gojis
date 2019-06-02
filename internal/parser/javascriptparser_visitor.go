// Code generated from JavaScriptParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // JavaScriptParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by JavaScriptParser.
type JavaScriptParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JavaScriptParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#sourceElement.
	VisitSourceElement(ctx *SourceElementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#variableStatement.
	VisitVariableStatement(ctx *VariableStatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#variableDeclarationList.
	VisitVariableDeclarationList(ctx *VariableDeclarationListContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#emptyStatement_.
	VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{}

	// Visit a parse tree produced by JavaScriptParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#DoStatement.
	VisitDoStatement(ctx *DoStatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#WhileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#ForStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#ForVarStatement.
	VisitForVarStatement(ctx *ForVarStatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#ForInStatement.
	VisitForInStatement(ctx *ForInStatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#ForVarInStatement.
	VisitForVarInStatement(ctx *ForVarInStatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#varModifier.
	VisitVarModifier(ctx *VarModifierContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#withStatement.
	VisitWithStatement(ctx *WithStatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#caseBlock.
	VisitCaseBlock(ctx *CaseBlockContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#caseClauses.
	VisitCaseClauses(ctx *CaseClausesContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#caseClause.
	VisitCaseClause(ctx *CaseClauseContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#defaultClause.
	VisitDefaultClause(ctx *DefaultClauseContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#labelledStatement.
	VisitLabelledStatement(ctx *LabelledStatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#throwStatement.
	VisitThrowStatement(ctx *ThrowStatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#tryStatement.
	VisitTryStatement(ctx *TryStatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#catchProduction.
	VisitCatchProduction(ctx *CatchProductionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#finallyProduction.
	VisitFinallyProduction(ctx *FinallyProductionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#debuggerStatement.
	VisitDebuggerStatement(ctx *DebuggerStatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#classTail.
	VisitClassTail(ctx *ClassTailContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#classElement.
	VisitClassElement(ctx *ClassElementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#methodDefinition.
	VisitMethodDefinition(ctx *MethodDefinitionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#generatorMethod.
	VisitGeneratorMethod(ctx *GeneratorMethodContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#formalParameterArg.
	VisitFormalParameterArg(ctx *FormalParameterArgContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#lastFormalParameterArg.
	VisitLastFormalParameterArg(ctx *LastFormalParameterArgContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#functionBody.
	VisitFunctionBody(ctx *FunctionBodyContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#sourceElements.
	VisitSourceElements(ctx *SourceElementsContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#elementList.
	VisitElementList(ctx *ElementListContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#lastElement.
	VisitLastElement(ctx *LastElementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#objectLiteral.
	VisitObjectLiteral(ctx *ObjectLiteralContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#PropertyExpressionAssignment.
	VisitPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#ComputedPropertyExpressionAssignment.
	VisitComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#PropertyGetter.
	VisitPropertyGetter(ctx *PropertyGetterContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#PropertySetter.
	VisitPropertySetter(ctx *PropertySetterContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#MethodProperty.
	VisitMethodProperty(ctx *MethodPropertyContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#PropertyShorthand.
	VisitPropertyShorthand(ctx *PropertyShorthandContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#propertyName.
	VisitPropertyName(ctx *PropertyNameContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#lastArgument.
	VisitLastArgument(ctx *LastArgumentContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#expressionSequence.
	VisitExpressionSequence(ctx *ExpressionSequenceContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#TemplateStringExpression.
	VisitTemplateStringExpression(ctx *TemplateStringExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#TernaryExpression.
	VisitTernaryExpression(ctx *TernaryExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#LogicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#PreIncrementExpression.
	VisitPreIncrementExpression(ctx *PreIncrementExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#ObjectLiteralExpression.
	VisitObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#InExpression.
	VisitInExpression(ctx *InExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#LogicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#NotExpression.
	VisitNotExpression(ctx *NotExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#PreDecreaseExpression.
	VisitPreDecreaseExpression(ctx *PreDecreaseExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#ArgumentsExpression.
	VisitArgumentsExpression(ctx *ArgumentsExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#ThisExpression.
	VisitThisExpression(ctx *ThisExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#FunctionExpression.
	VisitFunctionExpression(ctx *FunctionExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#UnaryMinusExpression.
	VisitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#AssignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#PostDecreaseExpression.
	VisitPostDecreaseExpression(ctx *PostDecreaseExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#TypeofExpression.
	VisitTypeofExpression(ctx *TypeofExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#InstanceofExpression.
	VisitInstanceofExpression(ctx *InstanceofExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#UnaryPlusExpression.
	VisitUnaryPlusExpression(ctx *UnaryPlusExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#DeleteExpression.
	VisitDeleteExpression(ctx *DeleteExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#ArrowFunctionExpression.
	VisitArrowFunctionExpression(ctx *ArrowFunctionExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#EqualityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#BitXOrExpression.
	VisitBitXOrExpression(ctx *BitXOrExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#SuperExpression.
	VisitSuperExpression(ctx *SuperExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#MultiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#BitShiftExpression.
	VisitBitShiftExpression(ctx *BitShiftExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#ParenthesizedExpression.
	VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#AdditiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#RelationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#PostIncrementExpression.
	VisitPostIncrementExpression(ctx *PostIncrementExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#BitNotExpression.
	VisitBitNotExpression(ctx *BitNotExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#NewExpression.
	VisitNewExpression(ctx *NewExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#LiteralExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#ArrayLiteralExpression.
	VisitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#MemberDotExpression.
	VisitMemberDotExpression(ctx *MemberDotExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#ClassExpression.
	VisitClassExpression(ctx *ClassExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#MemberIndexExpression.
	VisitMemberIndexExpression(ctx *MemberIndexExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#IdentifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#BitAndExpression.
	VisitBitAndExpression(ctx *BitAndExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#BitOrExpression.
	VisitBitOrExpression(ctx *BitOrExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#AssignmentOperatorExpression.
	VisitAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#VoidExpression.
	VisitVoidExpression(ctx *VoidExpressionContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#arrowFunctionParameters.
	VisitArrowFunctionParameters(ctx *ArrowFunctionParametersContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#arrowFunctionBody.
	VisitArrowFunctionBody(ctx *ArrowFunctionBodyContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#numericLiteral.
	VisitNumericLiteral(ctx *NumericLiteralContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#identifierName.
	VisitIdentifierName(ctx *IdentifierNameContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#reservedWord.
	VisitReservedWord(ctx *ReservedWordContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#getter.
	VisitGetter(ctx *GetterContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#setter.
	VisitSetter(ctx *SetterContext) interface{}

	// Visit a parse tree produced by JavaScriptParser#eos.
	VisitEos(ctx *EosContext) interface{}
}
