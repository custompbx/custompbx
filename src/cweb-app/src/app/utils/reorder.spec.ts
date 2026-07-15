import {reorderImmutable, resolvePositionedReorder} from './reorder';

describe('reorder utilities', () => {
  const items = [
    {id: 10, position: 1, name: 'one'},
    {id: 20, position: 4, name: 'two'},
    {id: 30, position: 9, name: 'three'},
  ];

  it('moves an item without mutating the source array', () => {
    const result = reorderImmutable(items, 0, 2);

    expect(result.map(item => item.id)).toEqual([20, 30, 10]);
    expect(items.map(item => item.id)).toEqual([10, 20, 30]);
  });

  it('returns the original array for unchanged or invalid drops', () => {
    expect(reorderImmutable(items, 1, 1)).toBe(items);
    expect(reorderImmutable(items, -1, 1)).toBe(items);
    expect(reorderImmutable(items, 0, 3)).toBe(items);
  });

  it('creates the existing backend move payload and stable local positions', () => {
    const result = resolvePositionedReorder(items, 2, 0);

    expect(result?.move).toEqual({previous_index: 9, current_index: 1, id: 30});
    expect(result?.items.map(item => [item.id, item.position])).toEqual([
      [30, 1], [10, 4], [20, 9],
    ]);
    expect(items[2].position).toBe(9);
  });

  it('rejects duplicate positions and out-of-bounds nested-list drops', () => {
    expect(resolvePositionedReorder([{id: 1, position: 2}, {id: 2, position: 2}], 0, 1)).toBeNull();
    expect(resolvePositionedReorder(items, 3, 0)).toBeNull();
  });
});
