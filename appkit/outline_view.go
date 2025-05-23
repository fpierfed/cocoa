// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var OutlineViewClass = _OutlineViewClass{objc.GetClass("NSOutlineView")}

type _OutlineViewClass struct {
	objc.Class
}

type IOutlineView interface {
	ITableView
	IsExpandable(item objc.IObject) bool
	IsItemExpanded(item objc.IObject) bool
	ExpandItem(item objc.IObject)
	ExpandItem_ExpandChildren(item objc.IObject, expandChildren bool)
	CollapseItem(item objc.IObject)
	CollapseItem_CollapseChildren(item objc.IObject, collapseChildren bool)
	ReloadItem(item objc.IObject)
	ReloadItem_ReloadChildren(item objc.IObject, reloadChildren bool)
	ItemAtRow(row int) objc.Object
	RowForItem(item objc.IObject) int
	LevelForItem(item objc.IObject) int
	LevelForRow(row int) int
	SetDropItem_DropChildIndex(item objc.IObject, index int)
	ShouldCollapseAutoExpandedItemsForDeposited(deposited bool) bool
	ParentForItem(item objc.IObject) objc.Object
	ChildIndexForItem(item objc.IObject) int
	Child_OfItem(index int, item objc.IObject) objc.Object
	NumberOfChildrenOfItem(item objc.IObject) int
	FrameOfOutlineCellAtRow(row int) foundation.Rect
	InsertItemsAtIndexes_InParent_WithAnimation(indexes foundation.IIndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions)
	MoveItemAtIndex_InParent_ToIndex_InParent(fromIndex int, oldParent objc.IObject, toIndex int, newParent objc.IObject)
	RemoveItemsAtIndexes_InParent_WithAnimation(indexes foundation.IIndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions)
	StronglyReferencesItems() bool
	SetStronglyReferencesItems(value bool)
	OutlineTableColumn() TableColumn
	SetOutlineTableColumn(value ITableColumn)
	AutoresizesOutlineColumn() bool
	SetAutoresizesOutlineColumn(value bool)
	IndentationPerLevel() float64
	SetIndentationPerLevel(value float64)
	IndentationMarkerFollowsCell() bool
	SetIndentationMarkerFollowsCell(value bool)
	AutosaveExpandedItems() bool
	SetAutosaveExpandedItems(value bool)
}

type OutlineView struct {
	TableView
}

func MakeOutlineView(ptr unsafe.Pointer) OutlineView {
	return OutlineView{
		TableView: MakeTableView(ptr),
	}
}

func (o_ OutlineView) InitWithFrame(frameRect foundation.Rect) OutlineView {
	rv := ffi.CallMethod[OutlineView](o_, "initWithFrame:", frameRect)
	return rv
}

func (o_ OutlineView) Init() OutlineView {
	rv := ffi.CallMethod[OutlineView](o_, "init")
	return rv
}

func (oc _OutlineViewClass) Alloc() OutlineView {
	rv := ffi.CallMethod[OutlineView](oc, "alloc")
	return rv
}

func (oc _OutlineViewClass) New() OutlineView {
	rv := ffi.CallMethod[OutlineView](oc, "new")
	rv.Autorelease()
	return rv
}

func NewOutlineView() OutlineView {
	return OutlineViewClass.New()
}

func (o_ OutlineView) IsExpandable(item objc.IObject) bool {
	rv := ffi.CallMethod[bool](o_, "isExpandable:", item)
	return rv
}

func (o_ OutlineView) IsItemExpanded(item objc.IObject) bool {
	rv := ffi.CallMethod[bool](o_, "isItemExpanded:", item)
	return rv
}

func (o_ OutlineView) ExpandItem(item objc.IObject) {
	ffi.CallMethod[ffi.Void](o_, "expandItem:", item)
}

func (o_ OutlineView) ExpandItem_ExpandChildren(item objc.IObject, expandChildren bool) {
	ffi.CallMethod[ffi.Void](o_, "expandItem:expandChildren:", item, expandChildren)
}

func (o_ OutlineView) CollapseItem(item objc.IObject) {
	ffi.CallMethod[ffi.Void](o_, "collapseItem:", item)
}

func (o_ OutlineView) CollapseItem_CollapseChildren(item objc.IObject, collapseChildren bool) {
	ffi.CallMethod[ffi.Void](o_, "collapseItem:collapseChildren:", item, collapseChildren)
}

func (o_ OutlineView) ReloadItem(item objc.IObject) {
	ffi.CallMethod[ffi.Void](o_, "reloadItem:", item)
}

func (o_ OutlineView) ReloadItem_ReloadChildren(item objc.IObject, reloadChildren bool) {
	ffi.CallMethod[ffi.Void](o_, "reloadItem:reloadChildren:", item, reloadChildren)
}

func (o_ OutlineView) ItemAtRow(row int) objc.Object {
	rv := ffi.CallMethod[objc.Object](o_, "itemAtRow:", row)
	return rv
}

func (o_ OutlineView) RowForItem(item objc.IObject) int {
	rv := ffi.CallMethod[int](o_, "rowForItem:", item)
	return rv
}

func (o_ OutlineView) LevelForItem(item objc.IObject) int {
	rv := ffi.CallMethod[int](o_, "levelForItem:", item)
	return rv
}

func (o_ OutlineView) LevelForRow(row int) int {
	rv := ffi.CallMethod[int](o_, "levelForRow:", row)
	return rv
}

func (o_ OutlineView) SetDropItem_DropChildIndex(item objc.IObject, index int) {
	ffi.CallMethod[ffi.Void](o_, "setDropItem:dropChildIndex:", item, index)
}

func (o_ OutlineView) ShouldCollapseAutoExpandedItemsForDeposited(deposited bool) bool {
	rv := ffi.CallMethod[bool](o_, "shouldCollapseAutoExpandedItemsForDeposited:", deposited)
	return rv
}

func (o_ OutlineView) ParentForItem(item objc.IObject) objc.Object {
	rv := ffi.CallMethod[objc.Object](o_, "parentForItem:", item)
	return rv
}

func (o_ OutlineView) ChildIndexForItem(item objc.IObject) int {
	rv := ffi.CallMethod[int](o_, "childIndexForItem:", item)
	return rv
}

func (o_ OutlineView) Child_OfItem(index int, item objc.IObject) objc.Object {
	rv := ffi.CallMethod[objc.Object](o_, "child:ofItem:", index, item)
	return rv
}

func (o_ OutlineView) NumberOfChildrenOfItem(item objc.IObject) int {
	rv := ffi.CallMethod[int](o_, "numberOfChildrenOfItem:", item)
	return rv
}

func (o_ OutlineView) FrameOfOutlineCellAtRow(row int) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](o_, "frameOfOutlineCellAtRow:", row)
	return rv
}

func (o_ OutlineView) InsertItemsAtIndexes_InParent_WithAnimation(indexes foundation.IIndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions) {
	ffi.CallMethod[ffi.Void](o_, "insertItemsAtIndexes:inParent:withAnimation:", indexes, parent, animationOptions)
}

func (o_ OutlineView) MoveItemAtIndex_InParent_ToIndex_InParent(fromIndex int, oldParent objc.IObject, toIndex int, newParent objc.IObject) {
	ffi.CallMethod[ffi.Void](o_, "moveItemAtIndex:inParent:toIndex:inParent:", fromIndex, oldParent, toIndex, newParent)
}

func (o_ OutlineView) RemoveItemsAtIndexes_InParent_WithAnimation(indexes foundation.IIndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions) {
	ffi.CallMethod[ffi.Void](o_, "removeItemsAtIndexes:inParent:withAnimation:", indexes, parent, animationOptions)
}

func (o_ OutlineView) StronglyReferencesItems() bool {
	rv := ffi.CallMethod[bool](o_, "stronglyReferencesItems")
	return rv
}

func (o_ OutlineView) SetStronglyReferencesItems(value bool) {
	ffi.CallMethod[ffi.Void](o_, "setStronglyReferencesItems:", value)
}

func (o_ OutlineView) OutlineTableColumn() TableColumn {
	rv := ffi.CallMethod[TableColumn](o_, "outlineTableColumn")
	return rv
}

func (o_ OutlineView) SetOutlineTableColumn(value ITableColumn) {
	ffi.CallMethod[ffi.Void](o_, "setOutlineTableColumn:", value)
}

func (o_ OutlineView) AutoresizesOutlineColumn() bool {
	rv := ffi.CallMethod[bool](o_, "autoresizesOutlineColumn")
	return rv
}

func (o_ OutlineView) SetAutoresizesOutlineColumn(value bool) {
	ffi.CallMethod[ffi.Void](o_, "setAutoresizesOutlineColumn:", value)
}

func (o_ OutlineView) IndentationPerLevel() float64 {
	rv := ffi.CallMethod[float64](o_, "indentationPerLevel")
	return rv
}

func (o_ OutlineView) SetIndentationPerLevel(value float64) {
	ffi.CallMethod[ffi.Void](o_, "setIndentationPerLevel:", value)
}

func (o_ OutlineView) IndentationMarkerFollowsCell() bool {
	rv := ffi.CallMethod[bool](o_, "indentationMarkerFollowsCell")
	return rv
}

func (o_ OutlineView) SetIndentationMarkerFollowsCell(value bool) {
	ffi.CallMethod[ffi.Void](o_, "setIndentationMarkerFollowsCell:", value)
}

func (o_ OutlineView) AutosaveExpandedItems() bool {
	rv := ffi.CallMethod[bool](o_, "autosaveExpandedItems")
	return rv
}

func (o_ OutlineView) SetAutosaveExpandedItems(value bool) {
	ffi.CallMethod[ffi.Void](o_, "setAutosaveExpandedItems:", value)
}

type OutlineViewDataSource interface {
	ImplementsOutlineView_AcceptDrop_Item_ChildIndex() bool
	// optional
	OutlineView_AcceptDrop_Item_ChildIndex(outlineView OutlineView, info DraggingInfoWrapper, item objc.Object, index int) bool
	ImplementsOutlineView_Child_OfItem() bool
	// optional
	OutlineView_Child_OfItem(outlineView OutlineView, index int, item objc.Object) objc.IObject
	ImplementsOutlineView_DraggingSession_EndedAtPoint_Operation() bool
	// optional
	OutlineView_DraggingSession_EndedAtPoint_Operation(outlineView OutlineView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	ImplementsOutlineView_DraggingSession_WillBeginAtPoint_ForItems() bool
	// optional
	OutlineView_DraggingSession_WillBeginAtPoint_ForItems(outlineView OutlineView, session DraggingSession, screenPoint foundation.Point, draggedItems []objc.Object)
	ImplementsOutlineView_IsItemExpandable() bool
	// optional
	OutlineView_IsItemExpandable(outlineView OutlineView, item objc.Object) bool
	ImplementsOutlineView_ItemForPersistentObject() bool
	// optional
	OutlineView_ItemForPersistentObject(outlineView OutlineView, object objc.Object) objc.IObject
	ImplementsOutlineView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItems() bool
	// optional
	// deprecated
	OutlineView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItems(outlineView OutlineView, dropDestination foundation.URL, items []objc.Object) []string
	ImplementsOutlineView_NumberOfChildrenOfItem() bool
	// optional
	OutlineView_NumberOfChildrenOfItem(outlineView OutlineView, item objc.Object) int
	ImplementsOutlineView_ObjectValueForTableColumn_ByItem() bool
	// optional
	OutlineView_ObjectValueForTableColumn_ByItem(outlineView OutlineView, tableColumn TableColumn, item objc.Object) objc.IObject
	ImplementsOutlineView_PasteboardWriterForItem() bool
	// optional
	OutlineView_PasteboardWriterForItem(outlineView OutlineView, item objc.Object) PasteboardWriting
	ImplementsOutlineView_PersistentObjectForItem() bool
	// optional
	OutlineView_PersistentObjectForItem(outlineView OutlineView, item objc.Object) objc.IObject
	ImplementsOutlineView_SetObjectValue_ForTableColumn_ByItem() bool
	// optional
	OutlineView_SetObjectValue_ForTableColumn_ByItem(outlineView OutlineView, object objc.Object, tableColumn TableColumn, item objc.Object)
	ImplementsOutlineView_SortDescriptorsDidChange() bool
	// optional
	OutlineView_SortDescriptorsDidChange(outlineView OutlineView, oldDescriptors []foundation.SortDescriptor)
	ImplementsOutlineView_UpdateDraggingItemsForDrag() bool
	// optional
	OutlineView_UpdateDraggingItemsForDrag(outlineView OutlineView, draggingInfo DraggingInfoWrapper)
	ImplementsOutlineView_ValidateDrop_ProposedItem_ProposedChildIndex() bool
	// optional
	OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex(outlineView OutlineView, info DraggingInfoWrapper, item objc.Object, index int) DragOperation
	ImplementsOutlineView_WriteItems_ToPasteboard() bool
	// optional
	// deprecated
	OutlineView_WriteItems_ToPasteboard(outlineView OutlineView, items []objc.Object, pasteboard Pasteboard) bool
}

type OutlineViewDataSourceWrapper struct {
	objc.Object
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_AcceptDrop_Item_ChildIndex() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:acceptDrop:item:childIndex:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_AcceptDrop_Item_ChildIndex(outlineView IOutlineView, info DraggingInfo, item objc.IObject, index int) bool {
	po := ffi.CreateProtocol(info)
	defer po.Release()
	rv := ffi.CallMethod[bool](o_, "outlineView:acceptDrop:item:childIndex:", outlineView, po, item, index)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_Child_OfItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:child:ofItem:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_Child_OfItem(outlineView IOutlineView, index int, item objc.IObject) objc.Object {
	rv := ffi.CallMethod[objc.Object](o_, "outlineView:child:ofItem:", outlineView, index, item)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_DraggingSession_EndedAtPoint_Operation() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:draggingSession:endedAtPoint:operation:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_DraggingSession_EndedAtPoint_Operation(outlineView IOutlineView, session IDraggingSession, screenPoint foundation.Point, operation DragOperation) {
	ffi.CallMethod[ffi.Void](o_, "outlineView:draggingSession:endedAtPoint:operation:", outlineView, session, screenPoint, operation)
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_DraggingSession_WillBeginAtPoint_ForItems() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:draggingSession:willBeginAtPoint:forItems:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_DraggingSession_WillBeginAtPoint_ForItems(outlineView IOutlineView, session IDraggingSession, screenPoint foundation.Point, draggedItems []objc.IObject) {
	ffi.CallMethod[ffi.Void](o_, "outlineView:draggingSession:willBeginAtPoint:forItems:", outlineView, session, screenPoint, draggedItems)
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_IsItemExpandable() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:isItemExpandable:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_IsItemExpandable(outlineView IOutlineView, item objc.IObject) bool {
	rv := ffi.CallMethod[bool](o_, "outlineView:isItemExpandable:", outlineView, item)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_ItemForPersistentObject() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:itemForPersistentObject:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_ItemForPersistentObject(outlineView IOutlineView, object objc.IObject) objc.Object {
	rv := ffi.CallMethod[objc.Object](o_, "outlineView:itemForPersistentObject:", outlineView, object)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItems() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:namesOfPromisedFilesDroppedAtDestination:forDraggedItems:"))
}

// deprecated
func (o_ OutlineViewDataSourceWrapper) OutlineView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItems(outlineView IOutlineView, dropDestination foundation.IURL, items []objc.IObject) []string {
	rv := ffi.CallMethod[[]string](o_, "outlineView:namesOfPromisedFilesDroppedAtDestination:forDraggedItems:", outlineView, dropDestination, items)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_NumberOfChildrenOfItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:numberOfChildrenOfItem:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_NumberOfChildrenOfItem(outlineView IOutlineView, item objc.IObject) int {
	rv := ffi.CallMethod[int](o_, "outlineView:numberOfChildrenOfItem:", outlineView, item)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_ObjectValueForTableColumn_ByItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:objectValueForTableColumn:byItem:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_ObjectValueForTableColumn_ByItem(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) objc.Object {
	rv := ffi.CallMethod[objc.Object](o_, "outlineView:objectValueForTableColumn:byItem:", outlineView, tableColumn, item)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_PasteboardWriterForItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:pasteboardWriterForItem:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_PasteboardWriterForItem(outlineView IOutlineView, item objc.IObject) PasteboardWritingWrapper {
	rv := ffi.CallMethod[PasteboardWritingWrapper](o_, "outlineView:pasteboardWriterForItem:", outlineView, item)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_PersistentObjectForItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:persistentObjectForItem:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_PersistentObjectForItem(outlineView IOutlineView, item objc.IObject) objc.Object {
	rv := ffi.CallMethod[objc.Object](o_, "outlineView:persistentObjectForItem:", outlineView, item)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_SetObjectValue_ForTableColumn_ByItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:setObjectValue:forTableColumn:byItem:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_SetObjectValue_ForTableColumn_ByItem(outlineView IOutlineView, object objc.IObject, tableColumn ITableColumn, item objc.IObject) {
	ffi.CallMethod[ffi.Void](o_, "outlineView:setObjectValue:forTableColumn:byItem:", outlineView, object, tableColumn, item)
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_SortDescriptorsDidChange() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:sortDescriptorsDidChange:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_SortDescriptorsDidChange(outlineView IOutlineView, oldDescriptors []foundation.ISortDescriptor) {
	ffi.CallMethod[ffi.Void](o_, "outlineView:sortDescriptorsDidChange:", outlineView, oldDescriptors)
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_UpdateDraggingItemsForDrag() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:updateDraggingItemsForDrag:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_UpdateDraggingItemsForDrag(outlineView IOutlineView, draggingInfo DraggingInfo) {
	po := ffi.CreateProtocol(draggingInfo)
	defer po.Release()
	ffi.CallMethod[ffi.Void](o_, "outlineView:updateDraggingItemsForDrag:", outlineView, po)
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_ValidateDrop_ProposedItem_ProposedChildIndex() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:validateDrop:proposedItem:proposedChildIndex:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex(outlineView IOutlineView, info DraggingInfo, item objc.IObject, index int) DragOperation {
	po := ffi.CreateProtocol(info)
	defer po.Release()
	rv := ffi.CallMethod[DragOperation](o_, "outlineView:validateDrop:proposedItem:proposedChildIndex:", outlineView, po, item, index)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_WriteItems_ToPasteboard() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:writeItems:toPasteboard:"))
}

// deprecated
func (o_ OutlineViewDataSourceWrapper) OutlineView_WriteItems_ToPasteboard(outlineView IOutlineView, items []objc.IObject, pasteboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](o_, "outlineView:writeItems:toPasteboard:", outlineView, items, pasteboard)
	return rv
}

type OutlineViewDelegate interface {
	ControlTextEditingDelegate
	ImplementsOutlineView_ShouldExpandItem() bool
	// optional
	OutlineView_ShouldExpandItem(outlineView OutlineView, item objc.Object) bool
	ImplementsOutlineView_ShouldCollapseItem() bool
	// optional
	OutlineView_ShouldCollapseItem(outlineView OutlineView, item objc.Object) bool
	ImplementsOutlineView_TypeSelectStringForTableColumn_Item() bool
	// optional
	OutlineView_TypeSelectStringForTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) string
	ImplementsOutlineView_NextTypeSelectMatchFromItem_ToItem_ForString() bool
	// optional
	OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString(outlineView OutlineView, startItem objc.Object, endItem objc.Object, searchString string) objc.IObject
	ImplementsOutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString() bool
	// optional
	OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString(outlineView OutlineView, event Event, searchString string) bool
	ImplementsOutlineView_ToolTipForCell_Rect_TableColumn_Item_MouseLocation() bool
	// optional
	OutlineView_ToolTipForCell_Rect_TableColumn_Item_MouseLocation(outlineView OutlineView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, item objc.Object, mouseLocation foundation.Point) string
	ImplementsOutlineView_ShouldSelectTableColumn() bool
	// optional
	OutlineView_ShouldSelectTableColumn(outlineView OutlineView, tableColumn TableColumn) bool
	ImplementsOutlineView_ShouldSelectItem() bool
	// optional
	OutlineView_ShouldSelectItem(outlineView OutlineView, item objc.Object) bool
	ImplementsOutlineView_SelectionIndexesForProposedSelection() bool
	// optional
	OutlineView_SelectionIndexesForProposedSelection(outlineView OutlineView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet
	ImplementsSelectionShouldChangeInOutlineView() bool
	// optional
	SelectionShouldChangeInOutlineView(outlineView OutlineView) bool
	ImplementsOutlineViewSelectionIsChanging() bool
	// optional
	OutlineViewSelectionIsChanging(notification foundation.Notification)
	ImplementsOutlineViewSelectionDidChange() bool
	// optional
	OutlineViewSelectionDidChange(notification foundation.Notification)
	ImplementsOutlineView_WillDisplayCell_ForTableColumn_Item() bool
	// optional
	OutlineView_WillDisplayCell_ForTableColumn_Item(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object)
	ImplementsOutlineView_WillDisplayOutlineCell_ForTableColumn_Item() bool
	// optional
	OutlineView_WillDisplayOutlineCell_ForTableColumn_Item(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object)
	ImplementsOutlineView_DataCellForTableColumn_Item() bool
	// optional
	OutlineView_DataCellForTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) ICell
	ImplementsOutlineView_ShouldShowOutlineCellForItem() bool
	// optional
	OutlineView_ShouldShowOutlineCellForItem(outlineView OutlineView, item objc.Object) bool
	ImplementsOutlineView_ShouldShowCellExpansionForTableColumn_Item() bool
	// optional
	OutlineView_ShouldShowCellExpansionForTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool
	ImplementsOutlineView_ShouldReorderColumn_ToColumn() bool
	// optional
	OutlineView_ShouldReorderColumn_ToColumn(outlineView OutlineView, columnIndex int, newColumnIndex int) bool
	ImplementsOutlineViewColumnDidMove() bool
	// optional
	OutlineViewColumnDidMove(notification foundation.Notification)
	ImplementsOutlineViewColumnDidResize() bool
	// optional
	OutlineViewColumnDidResize(notification foundation.Notification)
	ImplementsOutlineViewItemWillExpand() bool
	// optional
	OutlineViewItemWillExpand(notification foundation.Notification)
	ImplementsOutlineViewItemDidExpand() bool
	// optional
	OutlineViewItemDidExpand(notification foundation.Notification)
	ImplementsOutlineViewItemWillCollapse() bool
	// optional
	OutlineViewItemWillCollapse(notification foundation.Notification)
	ImplementsOutlineViewItemDidCollapse() bool
	// optional
	OutlineViewItemDidCollapse(notification foundation.Notification)
	ImplementsOutlineView_ShouldEditTableColumn_Item() bool
	// optional
	OutlineView_ShouldEditTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool
	ImplementsOutlineView_MouseDownInHeaderOfTableColumn() bool
	// optional
	OutlineView_MouseDownInHeaderOfTableColumn(outlineView OutlineView, tableColumn TableColumn)
	ImplementsOutlineView_DidClickTableColumn() bool
	// optional
	OutlineView_DidClickTableColumn(outlineView OutlineView, tableColumn TableColumn)
	ImplementsOutlineView_DidDragTableColumn() bool
	// optional
	OutlineView_DidDragTableColumn(outlineView OutlineView, tableColumn TableColumn)
	ImplementsOutlineView_HeightOfRowByItem() bool
	// optional
	OutlineView_HeightOfRowByItem(outlineView OutlineView, item objc.Object) float64
	ImplementsOutlineView_SizeToFitWidthOfColumn() bool
	// optional
	OutlineView_SizeToFitWidthOfColumn(outlineView OutlineView, column int) float64
	ImplementsOutlineView_TintConfigurationForItem() bool
	// optional
	OutlineView_TintConfigurationForItem(outlineView OutlineView, item objc.Object) ITintConfiguration
	ImplementsOutlineView_ShouldTrackCell_ForTableColumn_Item() bool
	// optional
	OutlineView_ShouldTrackCell_ForTableColumn_Item(outlineView OutlineView, cell Cell, tableColumn TableColumn, item objc.Object) bool
	ImplementsOutlineView_IsGroupItem() bool
	// optional
	OutlineView_IsGroupItem(outlineView OutlineView, item objc.Object) bool
	ImplementsOutlineView_DidAddRowView_ForRow() bool
	// optional
	OutlineView_DidAddRowView_ForRow(outlineView OutlineView, rowView TableRowView, row int)
	ImplementsOutlineView_DidRemoveRowView_ForRow() bool
	// optional
	OutlineView_DidRemoveRowView_ForRow(outlineView OutlineView, rowView TableRowView, row int)
	ImplementsOutlineView_RowViewForItem() bool
	// optional
	OutlineView_RowViewForItem(outlineView OutlineView, item objc.Object) ITableRowView
	ImplementsOutlineView_ViewForTableColumn_Item() bool
	// optional
	OutlineView_ViewForTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) IView
}

type OutlineViewDelegateImpl struct {
	ControlTextEditingDelegateImpl
	_OutlineView_ShouldExpandItem                                   func(outlineView OutlineView, item objc.Object) bool
	_OutlineView_ShouldCollapseItem                                 func(outlineView OutlineView, item objc.Object) bool
	_OutlineView_TypeSelectStringForTableColumn_Item                func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) string
	_OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString       func(outlineView OutlineView, startItem objc.Object, endItem objc.Object, searchString string) objc.IObject
	_OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString   func(outlineView OutlineView, event Event, searchString string) bool
	_OutlineView_ToolTipForCell_Rect_TableColumn_Item_MouseLocation func(outlineView OutlineView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, item objc.Object, mouseLocation foundation.Point) string
	_OutlineView_ShouldSelectTableColumn                            func(outlineView OutlineView, tableColumn TableColumn) bool
	_OutlineView_ShouldSelectItem                                   func(outlineView OutlineView, item objc.Object) bool
	_OutlineView_SelectionIndexesForProposedSelection               func(outlineView OutlineView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet
	_SelectionShouldChangeInOutlineView                             func(outlineView OutlineView) bool
	_OutlineViewSelectionIsChanging                                 func(notification foundation.Notification)
	_OutlineViewSelectionDidChange                                  func(notification foundation.Notification)
	_OutlineView_WillDisplayCell_ForTableColumn_Item                func(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object)
	_OutlineView_WillDisplayOutlineCell_ForTableColumn_Item         func(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object)
	_OutlineView_DataCellForTableColumn_Item                        func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) ICell
	_OutlineView_ShouldShowOutlineCellForItem                       func(outlineView OutlineView, item objc.Object) bool
	_OutlineView_ShouldShowCellExpansionForTableColumn_Item         func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool
	_OutlineView_ShouldReorderColumn_ToColumn                       func(outlineView OutlineView, columnIndex int, newColumnIndex int) bool
	_OutlineViewColumnDidMove                                       func(notification foundation.Notification)
	_OutlineViewColumnDidResize                                     func(notification foundation.Notification)
	_OutlineViewItemWillExpand                                      func(notification foundation.Notification)
	_OutlineViewItemDidExpand                                       func(notification foundation.Notification)
	_OutlineViewItemWillCollapse                                    func(notification foundation.Notification)
	_OutlineViewItemDidCollapse                                     func(notification foundation.Notification)
	_OutlineView_ShouldEditTableColumn_Item                         func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool
	_OutlineView_MouseDownInHeaderOfTableColumn                     func(outlineView OutlineView, tableColumn TableColumn)
	_OutlineView_DidClickTableColumn                                func(outlineView OutlineView, tableColumn TableColumn)
	_OutlineView_DidDragTableColumn                                 func(outlineView OutlineView, tableColumn TableColumn)
	_OutlineView_HeightOfRowByItem                                  func(outlineView OutlineView, item objc.Object) float64
	_OutlineView_SizeToFitWidthOfColumn                             func(outlineView OutlineView, column int) float64
	_OutlineView_TintConfigurationForItem                           func(outlineView OutlineView, item objc.Object) ITintConfiguration
	_OutlineView_ShouldTrackCell_ForTableColumn_Item                func(outlineView OutlineView, cell Cell, tableColumn TableColumn, item objc.Object) bool
	_OutlineView_IsGroupItem                                        func(outlineView OutlineView, item objc.Object) bool
	_OutlineView_DidAddRowView_ForRow                               func(outlineView OutlineView, rowView TableRowView, row int)
	_OutlineView_DidRemoveRowView_ForRow                            func(outlineView OutlineView, rowView TableRowView, row int)
	_OutlineView_RowViewForItem                                     func(outlineView OutlineView, item objc.Object) ITableRowView
	_OutlineView_ViewForTableColumn_Item                            func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) IView
}

func (di *OutlineViewDelegateImpl) ImplementsOutlineView_ShouldExpandItem() bool {
	return di._OutlineView_ShouldExpandItem != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_ShouldExpandItem(f func(outlineView OutlineView, item objc.Object) bool) {
	di._OutlineView_ShouldExpandItem = f
}

func (di *OutlineViewDelegateImpl) OutlineView_ShouldExpandItem(outlineView OutlineView, item objc.Object) bool {
	return di._OutlineView_ShouldExpandItem(outlineView, item)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_ShouldCollapseItem() bool {
	return di._OutlineView_ShouldCollapseItem != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_ShouldCollapseItem(f func(outlineView OutlineView, item objc.Object) bool) {
	di._OutlineView_ShouldCollapseItem = f
}

func (di *OutlineViewDelegateImpl) OutlineView_ShouldCollapseItem(outlineView OutlineView, item objc.Object) bool {
	return di._OutlineView_ShouldCollapseItem(outlineView, item)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_TypeSelectStringForTableColumn_Item() bool {
	return di._OutlineView_TypeSelectStringForTableColumn_Item != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_TypeSelectStringForTableColumn_Item(f func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) string) {
	di._OutlineView_TypeSelectStringForTableColumn_Item = f
}

func (di *OutlineViewDelegateImpl) OutlineView_TypeSelectStringForTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) string {
	return di._OutlineView_TypeSelectStringForTableColumn_Item(outlineView, tableColumn, item)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_NextTypeSelectMatchFromItem_ToItem_ForString() bool {
	return di._OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_NextTypeSelectMatchFromItem_ToItem_ForString(f func(outlineView OutlineView, startItem objc.Object, endItem objc.Object, searchString string) objc.IObject) {
	di._OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString = f
}

func (di *OutlineViewDelegateImpl) OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString(outlineView OutlineView, startItem objc.Object, endItem objc.Object, searchString string) objc.IObject {
	return di._OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString(outlineView, startItem, endItem, searchString)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString() bool {
	return di._OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString(f func(outlineView OutlineView, event Event, searchString string) bool) {
	di._OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString = f
}

func (di *OutlineViewDelegateImpl) OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString(outlineView OutlineView, event Event, searchString string) bool {
	return di._OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString(outlineView, event, searchString)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_ToolTipForCell_Rect_TableColumn_Item_MouseLocation() bool {
	return di._OutlineView_ToolTipForCell_Rect_TableColumn_Item_MouseLocation != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_ToolTipForCell_Rect_TableColumn_Item_MouseLocation(f func(outlineView OutlineView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, item objc.Object, mouseLocation foundation.Point) string) {
	di._OutlineView_ToolTipForCell_Rect_TableColumn_Item_MouseLocation = f
}

func (di *OutlineViewDelegateImpl) OutlineView_ToolTipForCell_Rect_TableColumn_Item_MouseLocation(outlineView OutlineView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, item objc.Object, mouseLocation foundation.Point) string {
	return di._OutlineView_ToolTipForCell_Rect_TableColumn_Item_MouseLocation(outlineView, cell, rect, tableColumn, item, mouseLocation)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_ShouldSelectTableColumn() bool {
	return di._OutlineView_ShouldSelectTableColumn != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_ShouldSelectTableColumn(f func(outlineView OutlineView, tableColumn TableColumn) bool) {
	di._OutlineView_ShouldSelectTableColumn = f
}

func (di *OutlineViewDelegateImpl) OutlineView_ShouldSelectTableColumn(outlineView OutlineView, tableColumn TableColumn) bool {
	return di._OutlineView_ShouldSelectTableColumn(outlineView, tableColumn)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_ShouldSelectItem() bool {
	return di._OutlineView_ShouldSelectItem != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_ShouldSelectItem(f func(outlineView OutlineView, item objc.Object) bool) {
	di._OutlineView_ShouldSelectItem = f
}

func (di *OutlineViewDelegateImpl) OutlineView_ShouldSelectItem(outlineView OutlineView, item objc.Object) bool {
	return di._OutlineView_ShouldSelectItem(outlineView, item)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_SelectionIndexesForProposedSelection() bool {
	return di._OutlineView_SelectionIndexesForProposedSelection != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_SelectionIndexesForProposedSelection(f func(outlineView OutlineView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet) {
	di._OutlineView_SelectionIndexesForProposedSelection = f
}

func (di *OutlineViewDelegateImpl) OutlineView_SelectionIndexesForProposedSelection(outlineView OutlineView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet {
	return di._OutlineView_SelectionIndexesForProposedSelection(outlineView, proposedSelectionIndexes)
}
func (di *OutlineViewDelegateImpl) ImplementsSelectionShouldChangeInOutlineView() bool {
	return di._SelectionShouldChangeInOutlineView != nil
}

func (di *OutlineViewDelegateImpl) SetSelectionShouldChangeInOutlineView(f func(outlineView OutlineView) bool) {
	di._SelectionShouldChangeInOutlineView = f
}

func (di *OutlineViewDelegateImpl) SelectionShouldChangeInOutlineView(outlineView OutlineView) bool {
	return di._SelectionShouldChangeInOutlineView(outlineView)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineViewSelectionIsChanging() bool {
	return di._OutlineViewSelectionIsChanging != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineViewSelectionIsChanging(f func(notification foundation.Notification)) {
	di._OutlineViewSelectionIsChanging = f
}

func (di *OutlineViewDelegateImpl) OutlineViewSelectionIsChanging(notification foundation.Notification) {
	di._OutlineViewSelectionIsChanging(notification)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineViewSelectionDidChange() bool {
	return di._OutlineViewSelectionDidChange != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineViewSelectionDidChange(f func(notification foundation.Notification)) {
	di._OutlineViewSelectionDidChange = f
}

func (di *OutlineViewDelegateImpl) OutlineViewSelectionDidChange(notification foundation.Notification) {
	di._OutlineViewSelectionDidChange(notification)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_WillDisplayCell_ForTableColumn_Item() bool {
	return di._OutlineView_WillDisplayCell_ForTableColumn_Item != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_WillDisplayCell_ForTableColumn_Item(f func(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object)) {
	di._OutlineView_WillDisplayCell_ForTableColumn_Item = f
}

func (di *OutlineViewDelegateImpl) OutlineView_WillDisplayCell_ForTableColumn_Item(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object) {
	di._OutlineView_WillDisplayCell_ForTableColumn_Item(outlineView, cell, tableColumn, item)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_WillDisplayOutlineCell_ForTableColumn_Item() bool {
	return di._OutlineView_WillDisplayOutlineCell_ForTableColumn_Item != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_WillDisplayOutlineCell_ForTableColumn_Item(f func(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object)) {
	di._OutlineView_WillDisplayOutlineCell_ForTableColumn_Item = f
}

func (di *OutlineViewDelegateImpl) OutlineView_WillDisplayOutlineCell_ForTableColumn_Item(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object) {
	di._OutlineView_WillDisplayOutlineCell_ForTableColumn_Item(outlineView, cell, tableColumn, item)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_DataCellForTableColumn_Item() bool {
	return di._OutlineView_DataCellForTableColumn_Item != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_DataCellForTableColumn_Item(f func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) ICell) {
	di._OutlineView_DataCellForTableColumn_Item = f
}

func (di *OutlineViewDelegateImpl) OutlineView_DataCellForTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) ICell {
	return di._OutlineView_DataCellForTableColumn_Item(outlineView, tableColumn, item)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_ShouldShowOutlineCellForItem() bool {
	return di._OutlineView_ShouldShowOutlineCellForItem != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_ShouldShowOutlineCellForItem(f func(outlineView OutlineView, item objc.Object) bool) {
	di._OutlineView_ShouldShowOutlineCellForItem = f
}

func (di *OutlineViewDelegateImpl) OutlineView_ShouldShowOutlineCellForItem(outlineView OutlineView, item objc.Object) bool {
	return di._OutlineView_ShouldShowOutlineCellForItem(outlineView, item)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_ShouldShowCellExpansionForTableColumn_Item() bool {
	return di._OutlineView_ShouldShowCellExpansionForTableColumn_Item != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_ShouldShowCellExpansionForTableColumn_Item(f func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool) {
	di._OutlineView_ShouldShowCellExpansionForTableColumn_Item = f
}

func (di *OutlineViewDelegateImpl) OutlineView_ShouldShowCellExpansionForTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool {
	return di._OutlineView_ShouldShowCellExpansionForTableColumn_Item(outlineView, tableColumn, item)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_ShouldReorderColumn_ToColumn() bool {
	return di._OutlineView_ShouldReorderColumn_ToColumn != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_ShouldReorderColumn_ToColumn(f func(outlineView OutlineView, columnIndex int, newColumnIndex int) bool) {
	di._OutlineView_ShouldReorderColumn_ToColumn = f
}

func (di *OutlineViewDelegateImpl) OutlineView_ShouldReorderColumn_ToColumn(outlineView OutlineView, columnIndex int, newColumnIndex int) bool {
	return di._OutlineView_ShouldReorderColumn_ToColumn(outlineView, columnIndex, newColumnIndex)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineViewColumnDidMove() bool {
	return di._OutlineViewColumnDidMove != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineViewColumnDidMove(f func(notification foundation.Notification)) {
	di._OutlineViewColumnDidMove = f
}

func (di *OutlineViewDelegateImpl) OutlineViewColumnDidMove(notification foundation.Notification) {
	di._OutlineViewColumnDidMove(notification)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineViewColumnDidResize() bool {
	return di._OutlineViewColumnDidResize != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineViewColumnDidResize(f func(notification foundation.Notification)) {
	di._OutlineViewColumnDidResize = f
}

func (di *OutlineViewDelegateImpl) OutlineViewColumnDidResize(notification foundation.Notification) {
	di._OutlineViewColumnDidResize(notification)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineViewItemWillExpand() bool {
	return di._OutlineViewItemWillExpand != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineViewItemWillExpand(f func(notification foundation.Notification)) {
	di._OutlineViewItemWillExpand = f
}

func (di *OutlineViewDelegateImpl) OutlineViewItemWillExpand(notification foundation.Notification) {
	di._OutlineViewItemWillExpand(notification)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineViewItemDidExpand() bool {
	return di._OutlineViewItemDidExpand != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineViewItemDidExpand(f func(notification foundation.Notification)) {
	di._OutlineViewItemDidExpand = f
}

func (di *OutlineViewDelegateImpl) OutlineViewItemDidExpand(notification foundation.Notification) {
	di._OutlineViewItemDidExpand(notification)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineViewItemWillCollapse() bool {
	return di._OutlineViewItemWillCollapse != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineViewItemWillCollapse(f func(notification foundation.Notification)) {
	di._OutlineViewItemWillCollapse = f
}

func (di *OutlineViewDelegateImpl) OutlineViewItemWillCollapse(notification foundation.Notification) {
	di._OutlineViewItemWillCollapse(notification)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineViewItemDidCollapse() bool {
	return di._OutlineViewItemDidCollapse != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineViewItemDidCollapse(f func(notification foundation.Notification)) {
	di._OutlineViewItemDidCollapse = f
}

func (di *OutlineViewDelegateImpl) OutlineViewItemDidCollapse(notification foundation.Notification) {
	di._OutlineViewItemDidCollapse(notification)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_ShouldEditTableColumn_Item() bool {
	return di._OutlineView_ShouldEditTableColumn_Item != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_ShouldEditTableColumn_Item(f func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool) {
	di._OutlineView_ShouldEditTableColumn_Item = f
}

func (di *OutlineViewDelegateImpl) OutlineView_ShouldEditTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool {
	return di._OutlineView_ShouldEditTableColumn_Item(outlineView, tableColumn, item)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_MouseDownInHeaderOfTableColumn() bool {
	return di._OutlineView_MouseDownInHeaderOfTableColumn != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_MouseDownInHeaderOfTableColumn(f func(outlineView OutlineView, tableColumn TableColumn)) {
	di._OutlineView_MouseDownInHeaderOfTableColumn = f
}

func (di *OutlineViewDelegateImpl) OutlineView_MouseDownInHeaderOfTableColumn(outlineView OutlineView, tableColumn TableColumn) {
	di._OutlineView_MouseDownInHeaderOfTableColumn(outlineView, tableColumn)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_DidClickTableColumn() bool {
	return di._OutlineView_DidClickTableColumn != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_DidClickTableColumn(f func(outlineView OutlineView, tableColumn TableColumn)) {
	di._OutlineView_DidClickTableColumn = f
}

func (di *OutlineViewDelegateImpl) OutlineView_DidClickTableColumn(outlineView OutlineView, tableColumn TableColumn) {
	di._OutlineView_DidClickTableColumn(outlineView, tableColumn)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_DidDragTableColumn() bool {
	return di._OutlineView_DidDragTableColumn != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_DidDragTableColumn(f func(outlineView OutlineView, tableColumn TableColumn)) {
	di._OutlineView_DidDragTableColumn = f
}

func (di *OutlineViewDelegateImpl) OutlineView_DidDragTableColumn(outlineView OutlineView, tableColumn TableColumn) {
	di._OutlineView_DidDragTableColumn(outlineView, tableColumn)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_HeightOfRowByItem() bool {
	return di._OutlineView_HeightOfRowByItem != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_HeightOfRowByItem(f func(outlineView OutlineView, item objc.Object) float64) {
	di._OutlineView_HeightOfRowByItem = f
}

func (di *OutlineViewDelegateImpl) OutlineView_HeightOfRowByItem(outlineView OutlineView, item objc.Object) float64 {
	return di._OutlineView_HeightOfRowByItem(outlineView, item)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_SizeToFitWidthOfColumn() bool {
	return di._OutlineView_SizeToFitWidthOfColumn != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_SizeToFitWidthOfColumn(f func(outlineView OutlineView, column int) float64) {
	di._OutlineView_SizeToFitWidthOfColumn = f
}

func (di *OutlineViewDelegateImpl) OutlineView_SizeToFitWidthOfColumn(outlineView OutlineView, column int) float64 {
	return di._OutlineView_SizeToFitWidthOfColumn(outlineView, column)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_TintConfigurationForItem() bool {
	return di._OutlineView_TintConfigurationForItem != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_TintConfigurationForItem(f func(outlineView OutlineView, item objc.Object) ITintConfiguration) {
	di._OutlineView_TintConfigurationForItem = f
}

func (di *OutlineViewDelegateImpl) OutlineView_TintConfigurationForItem(outlineView OutlineView, item objc.Object) ITintConfiguration {
	return di._OutlineView_TintConfigurationForItem(outlineView, item)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_ShouldTrackCell_ForTableColumn_Item() bool {
	return di._OutlineView_ShouldTrackCell_ForTableColumn_Item != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_ShouldTrackCell_ForTableColumn_Item(f func(outlineView OutlineView, cell Cell, tableColumn TableColumn, item objc.Object) bool) {
	di._OutlineView_ShouldTrackCell_ForTableColumn_Item = f
}

func (di *OutlineViewDelegateImpl) OutlineView_ShouldTrackCell_ForTableColumn_Item(outlineView OutlineView, cell Cell, tableColumn TableColumn, item objc.Object) bool {
	return di._OutlineView_ShouldTrackCell_ForTableColumn_Item(outlineView, cell, tableColumn, item)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_IsGroupItem() bool {
	return di._OutlineView_IsGroupItem != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_IsGroupItem(f func(outlineView OutlineView, item objc.Object) bool) {
	di._OutlineView_IsGroupItem = f
}

func (di *OutlineViewDelegateImpl) OutlineView_IsGroupItem(outlineView OutlineView, item objc.Object) bool {
	return di._OutlineView_IsGroupItem(outlineView, item)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_DidAddRowView_ForRow() bool {
	return di._OutlineView_DidAddRowView_ForRow != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_DidAddRowView_ForRow(f func(outlineView OutlineView, rowView TableRowView, row int)) {
	di._OutlineView_DidAddRowView_ForRow = f
}

func (di *OutlineViewDelegateImpl) OutlineView_DidAddRowView_ForRow(outlineView OutlineView, rowView TableRowView, row int) {
	di._OutlineView_DidAddRowView_ForRow(outlineView, rowView, row)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_DidRemoveRowView_ForRow() bool {
	return di._OutlineView_DidRemoveRowView_ForRow != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_DidRemoveRowView_ForRow(f func(outlineView OutlineView, rowView TableRowView, row int)) {
	di._OutlineView_DidRemoveRowView_ForRow = f
}

func (di *OutlineViewDelegateImpl) OutlineView_DidRemoveRowView_ForRow(outlineView OutlineView, rowView TableRowView, row int) {
	di._OutlineView_DidRemoveRowView_ForRow(outlineView, rowView, row)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_RowViewForItem() bool {
	return di._OutlineView_RowViewForItem != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_RowViewForItem(f func(outlineView OutlineView, item objc.Object) ITableRowView) {
	di._OutlineView_RowViewForItem = f
}

func (di *OutlineViewDelegateImpl) OutlineView_RowViewForItem(outlineView OutlineView, item objc.Object) ITableRowView {
	return di._OutlineView_RowViewForItem(outlineView, item)
}
func (di *OutlineViewDelegateImpl) ImplementsOutlineView_ViewForTableColumn_Item() bool {
	return di._OutlineView_ViewForTableColumn_Item != nil
}

func (di *OutlineViewDelegateImpl) SetOutlineView_ViewForTableColumn_Item(f func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) IView) {
	di._OutlineView_ViewForTableColumn_Item = f
}

func (di *OutlineViewDelegateImpl) OutlineView_ViewForTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) IView {
	return di._OutlineView_ViewForTableColumn_Item(outlineView, tableColumn, item)
}

type OutlineViewDelegateWrapper struct {
	ControlTextEditingDelegateWrapper
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_ShouldExpandItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldExpandItem:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldExpandItem(outlineView IOutlineView, item objc.IObject) bool {
	rv := ffi.CallMethod[bool](o_, "outlineView:shouldExpandItem:", outlineView, item)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_ShouldCollapseItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldCollapseItem:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldCollapseItem(outlineView IOutlineView, item objc.IObject) bool {
	rv := ffi.CallMethod[bool](o_, "outlineView:shouldCollapseItem:", outlineView, item)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_TypeSelectStringForTableColumn_Item() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:typeSelectStringForTableColumn:item:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_TypeSelectStringForTableColumn_Item(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) string {
	rv := ffi.CallMethod[string](o_, "outlineView:typeSelectStringForTableColumn:item:", outlineView, tableColumn, item)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_NextTypeSelectMatchFromItem_ToItem_ForString() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:nextTypeSelectMatchFromItem:toItem:forString:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString(outlineView IOutlineView, startItem objc.IObject, endItem objc.IObject, searchString string) objc.Object {
	rv := ffi.CallMethod[objc.Object](o_, "outlineView:nextTypeSelectMatchFromItem:toItem:forString:", outlineView, startItem, endItem, searchString)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldTypeSelectForEvent:withCurrentSearchString:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString(outlineView IOutlineView, event IEvent, searchString string) bool {
	rv := ffi.CallMethod[bool](o_, "outlineView:shouldTypeSelectForEvent:withCurrentSearchString:", outlineView, event, searchString)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_ToolTipForCell_Rect_TableColumn_Item_MouseLocation() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:toolTipForCell:rect:tableColumn:item:mouseLocation:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ToolTipForCell_Rect_TableColumn_Item_MouseLocation(outlineView IOutlineView, cell ICell, rect *foundation.Rect, tableColumn ITableColumn, item objc.IObject, mouseLocation foundation.Point) string {
	rv := ffi.CallMethod[string](o_, "outlineView:toolTipForCell:rect:tableColumn:item:mouseLocation:", outlineView, cell, rect, tableColumn, item, mouseLocation)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_ShouldSelectTableColumn() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldSelectTableColumn:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldSelectTableColumn(outlineView IOutlineView, tableColumn ITableColumn) bool {
	rv := ffi.CallMethod[bool](o_, "outlineView:shouldSelectTableColumn:", outlineView, tableColumn)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_ShouldSelectItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldSelectItem:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldSelectItem(outlineView IOutlineView, item objc.IObject) bool {
	rv := ffi.CallMethod[bool](o_, "outlineView:shouldSelectItem:", outlineView, item)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_SelectionIndexesForProposedSelection() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:selectionIndexesForProposedSelection:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_SelectionIndexesForProposedSelection(outlineView IOutlineView, proposedSelectionIndexes foundation.IIndexSet) foundation.IndexSet {
	rv := ffi.CallMethod[foundation.IndexSet](o_, "outlineView:selectionIndexesForProposedSelection:", outlineView, proposedSelectionIndexes)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsSelectionShouldChangeInOutlineView() bool {
	return o_.RespondsToSelector(objc.GetSelector("selectionShouldChangeInOutlineView:"))
}

func (o_ OutlineViewDelegateWrapper) SelectionShouldChangeInOutlineView(outlineView IOutlineView) bool {
	rv := ffi.CallMethod[bool](o_, "selectionShouldChangeInOutlineView:", outlineView)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineViewSelectionIsChanging() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineViewSelectionIsChanging:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewSelectionIsChanging(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](o_, "outlineViewSelectionIsChanging:", notification)
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineViewSelectionDidChange() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineViewSelectionDidChange:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewSelectionDidChange(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](o_, "outlineViewSelectionDidChange:", notification)
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_WillDisplayCell_ForTableColumn_Item() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:willDisplayCell:forTableColumn:item:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_WillDisplayCell_ForTableColumn_Item(outlineView IOutlineView, cell objc.IObject, tableColumn ITableColumn, item objc.IObject) {
	ffi.CallMethod[ffi.Void](o_, "outlineView:willDisplayCell:forTableColumn:item:", outlineView, cell, tableColumn, item)
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_WillDisplayOutlineCell_ForTableColumn_Item() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:willDisplayOutlineCell:forTableColumn:item:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_WillDisplayOutlineCell_ForTableColumn_Item(outlineView IOutlineView, cell objc.IObject, tableColumn ITableColumn, item objc.IObject) {
	ffi.CallMethod[ffi.Void](o_, "outlineView:willDisplayOutlineCell:forTableColumn:item:", outlineView, cell, tableColumn, item)
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_DataCellForTableColumn_Item() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:dataCellForTableColumn:item:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_DataCellForTableColumn_Item(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) Cell {
	rv := ffi.CallMethod[Cell](o_, "outlineView:dataCellForTableColumn:item:", outlineView, tableColumn, item)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_ShouldShowOutlineCellForItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldShowOutlineCellForItem:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldShowOutlineCellForItem(outlineView IOutlineView, item objc.IObject) bool {
	rv := ffi.CallMethod[bool](o_, "outlineView:shouldShowOutlineCellForItem:", outlineView, item)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_ShouldShowCellExpansionForTableColumn_Item() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldShowCellExpansionForTableColumn:item:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldShowCellExpansionForTableColumn_Item(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) bool {
	rv := ffi.CallMethod[bool](o_, "outlineView:shouldShowCellExpansionForTableColumn:item:", outlineView, tableColumn, item)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_ShouldReorderColumn_ToColumn() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldReorderColumn:toColumn:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldReorderColumn_ToColumn(outlineView IOutlineView, columnIndex int, newColumnIndex int) bool {
	rv := ffi.CallMethod[bool](o_, "outlineView:shouldReorderColumn:toColumn:", outlineView, columnIndex, newColumnIndex)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineViewColumnDidMove() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineViewColumnDidMove:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewColumnDidMove(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](o_, "outlineViewColumnDidMove:", notification)
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineViewColumnDidResize() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineViewColumnDidResize:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewColumnDidResize(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](o_, "outlineViewColumnDidResize:", notification)
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineViewItemWillExpand() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineViewItemWillExpand:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewItemWillExpand(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](o_, "outlineViewItemWillExpand:", notification)
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineViewItemDidExpand() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineViewItemDidExpand:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewItemDidExpand(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](o_, "outlineViewItemDidExpand:", notification)
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineViewItemWillCollapse() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineViewItemWillCollapse:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewItemWillCollapse(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](o_, "outlineViewItemWillCollapse:", notification)
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineViewItemDidCollapse() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineViewItemDidCollapse:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewItemDidCollapse(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](o_, "outlineViewItemDidCollapse:", notification)
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_ShouldEditTableColumn_Item() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldEditTableColumn:item:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldEditTableColumn_Item(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) bool {
	rv := ffi.CallMethod[bool](o_, "outlineView:shouldEditTableColumn:item:", outlineView, tableColumn, item)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_MouseDownInHeaderOfTableColumn() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:mouseDownInHeaderOfTableColumn:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_MouseDownInHeaderOfTableColumn(outlineView IOutlineView, tableColumn ITableColumn) {
	ffi.CallMethod[ffi.Void](o_, "outlineView:mouseDownInHeaderOfTableColumn:", outlineView, tableColumn)
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_DidClickTableColumn() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:didClickTableColumn:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_DidClickTableColumn(outlineView IOutlineView, tableColumn ITableColumn) {
	ffi.CallMethod[ffi.Void](o_, "outlineView:didClickTableColumn:", outlineView, tableColumn)
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_DidDragTableColumn() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:didDragTableColumn:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_DidDragTableColumn(outlineView IOutlineView, tableColumn ITableColumn) {
	ffi.CallMethod[ffi.Void](o_, "outlineView:didDragTableColumn:", outlineView, tableColumn)
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_HeightOfRowByItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:heightOfRowByItem:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_HeightOfRowByItem(outlineView IOutlineView, item objc.IObject) float64 {
	rv := ffi.CallMethod[float64](o_, "outlineView:heightOfRowByItem:", outlineView, item)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_SizeToFitWidthOfColumn() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:sizeToFitWidthOfColumn:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_SizeToFitWidthOfColumn(outlineView IOutlineView, column int) float64 {
	rv := ffi.CallMethod[float64](o_, "outlineView:sizeToFitWidthOfColumn:", outlineView, column)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_TintConfigurationForItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:tintConfigurationForItem:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_TintConfigurationForItem(outlineView IOutlineView, item objc.IObject) TintConfiguration {
	rv := ffi.CallMethod[TintConfiguration](o_, "outlineView:tintConfigurationForItem:", outlineView, item)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_ShouldTrackCell_ForTableColumn_Item() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldTrackCell:forTableColumn:item:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldTrackCell_ForTableColumn_Item(outlineView IOutlineView, cell ICell, tableColumn ITableColumn, item objc.IObject) bool {
	rv := ffi.CallMethod[bool](o_, "outlineView:shouldTrackCell:forTableColumn:item:", outlineView, cell, tableColumn, item)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_IsGroupItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:isGroupItem:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_IsGroupItem(outlineView IOutlineView, item objc.IObject) bool {
	rv := ffi.CallMethod[bool](o_, "outlineView:isGroupItem:", outlineView, item)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_DidAddRowView_ForRow() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:didAddRowView:forRow:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_DidAddRowView_ForRow(outlineView IOutlineView, rowView ITableRowView, row int) {
	ffi.CallMethod[ffi.Void](o_, "outlineView:didAddRowView:forRow:", outlineView, rowView, row)
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_DidRemoveRowView_ForRow() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:didRemoveRowView:forRow:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_DidRemoveRowView_ForRow(outlineView IOutlineView, rowView ITableRowView, row int) {
	ffi.CallMethod[ffi.Void](o_, "outlineView:didRemoveRowView:forRow:", outlineView, rowView, row)
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_RowViewForItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:rowViewForItem:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_RowViewForItem(outlineView IOutlineView, item objc.IObject) TableRowView {
	rv := ffi.CallMethod[TableRowView](o_, "outlineView:rowViewForItem:", outlineView, item)
	return rv
}

func (o_ *OutlineViewDelegateWrapper) ImplementsOutlineView_ViewForTableColumn_Item() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:viewForTableColumn:item:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ViewForTableColumn_Item(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) View {
	rv := ffi.CallMethod[View](o_, "outlineView:viewForTableColumn:item:", outlineView, tableColumn, item)
	return rv
}
