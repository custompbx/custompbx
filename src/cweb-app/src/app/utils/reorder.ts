export interface PositionedItem {
  id: number;
  position: number;
}

export interface PositionMove {
  previous_index: number;
  current_index: number;
  id: number;
}

export interface PositionedReorder<T> {
  items: readonly T[];
  move: PositionMove;
}

function validIndex(items: readonly unknown[], index: number): boolean {
  return Number.isInteger(index) && index >= 0 && index < items.length;
}

export function reorderImmutable<T>(items: readonly T[], previousIndex: number, currentIndex: number): readonly T[] {
  if (!validIndex(items, previousIndex) || !validIndex(items, currentIndex) || previousIndex === currentIndex) {
    return items;
  }

  const reordered = [...items];
  const [moved] = reordered.splice(previousIndex, 1);
  reordered.splice(currentIndex, 0, moved);
  return reordered;
}

export function resolvePositionedReorder<T extends PositionedItem>(
  items: readonly T[],
  previousIndex: number,
  currentIndex: number,
): PositionedReorder<T> | null {
  if (!validIndex(items, previousIndex) || !validIndex(items, currentIndex) || previousIndex === currentIndex) {
    return null;
  }

  const moved = items[previousIndex];
  const target = items[currentIndex];
  if (!Number.isFinite(moved.position) || !Number.isFinite(target.position) || moved.position === target.position) {
    return null;
  }

  const positionSlots = items.map(item => item.position);
  const reordered = reorderImmutable(items, previousIndex, currentIndex).map((item, index) =>
    item.position === positionSlots[index] ? item : {...item, position: positionSlots[index]}
  );

  return {
    items: reordered,
    move: {
      previous_index: moved.position,
      current_index: target.position,
      id: moved.id,
    },
  };
}
