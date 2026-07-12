import {KeyValuePad2Component} from './key-value-pad-2.component';

describe('KeyValuePad2Component', () => {
  let component: KeyValuePad2Component;

  beforeEach(() => {
    component = new KeyValuePad2Component(null as any);
    component.ngOnInit();
  });

  it('uses compact field class for narrow configured fields', () => {
    const classes = component.fieldClass({
      slot: 'extraField1',
      name: 'secure',
      style: {'max-width': '96px'},
    } as any);

    expect(classes).toContain('kv-field--sm');
  });

  it('uses explicit field size metadata before legacy inline styles', () => {
    const classes = component.fieldClass({
      slot: 'extraField1',
      name: 'priority',
      size: 'sm',
      style: {'width': '520px'},
    } as any);

    expect(classes).toContain('kv-field--sm');
    expect(classes).not.toContain('kv-field--wide');
  });

  it('uses select field class when options are configured', () => {
    const classes = component.fieldClass({
      slot: 'extraField1',
      name: 'secure',
      options: ['', 'true', 'false'],
    } as any);

    expect(classes).toContain('kv-field--select');
  });

  it('keeps explicit required metadata for non-name fields', () => {
    component.fieldsMask = {
      name: {name: 'name'},
      value: {name: 'value', required: true},
      extraField1: {name: 'secure', required: false},
    };

    const configs = component.fieldConfigs();

    expect(configs.find((field) => field.slot === 'name')?.required).toBeTrue();
    expect(configs.find((field) => field.slot === 'value')?.required).toBeTrue();
    expect(configs.find((field) => field.slot === 'extraField1')?.required).toBeFalse();
  });

  it('labels empty select values as Default', () => {
    expect(component.optionLabel('')).toBe('Default');
    expect(component.optionLabel('true')).toBe('true');
  });

  it('sorts values by position only when sortable is enabled', () => {
    const values = {
      a: {id: 1, position: 2},
      b: {id: 2, position: 1},
      new: [],
    };

    expect(component.displayValues(values).map((item) => item.id)).toEqual([1, 2]);

    component.sortable = true;

    expect(component.displayValues(values).map((item) => item.id)).toEqual([2, 1]);
  });

  it('shows filter only when more than one real item exists', () => {
    expect(component.showFilter({a: {id: 1, name: 'only'}, new: []})).toBeFalse();
    expect(component.showFilter({a: {id: 1, name: 'one'}, b: {id: 2, name: 'two'}})).toBeTrue();
  });

  it('filters by configured name, value, extra field, or id', () => {
    component.fieldsMask = {
      name: {name: 'name'},
      value: {name: 'value'},
      extraField1: {name: 'secure'},
    };
    component.filterText = 'tls';

    const values = component.filteredValues({
      a: {id: 1, name: 'alpha', value: '', secure: 'tls'},
      b: {id: 2, name: 'beta', value: 'udp', secure: ''},
      new: [],
    });

    expect(values.map((item) => item.id)).toEqual([1]);
  });

  it('rejects blank and duplicate item names', () => {
    component.items = {
      a: {id: 1, name: 'existing'},
      b: {id: 2, name: 'other'},
    };
    component.newItems = [{name: 'draft'}];

    expect(component.nameValidationMessage('   ')).toBe('Name is required.');
    expect(component.nameValidationMessage('Existing', 2)).toBe('Another item already uses this name.');
    expect(component.nameValidationMessage('existing', 1)).toBeNull();
    expect(component.nameValidationMessage('draft', null, 1)).toBe('Another item already uses this name.');
  });

  it('trims and maps update payloads through configured field names', () => {
    const updateItem = jasmine.createSpy('updateItem');
    component.fieldsMask = {
      name: {name: 'name'},
      value: {name: 'value'},
      extraField1: {name: 'secure'},
    };
    component.dispatchersCallbacks = {updateItem};

    component.updateItem(7, '  password ', ' secret ', ' true ');

    expect(updateItem).toHaveBeenCalledOnceWith({
      id: 7,
      name: 'password',
      value: 'secret',
      secure: 'true',
    });
  });

  it('keeps empty values when adding an item but drops missing extra fields', () => {
    const addItem = jasmine.createSpy('addItem');
    component.id = 12;
    component.fieldsMask = {
      name: {name: 'name'},
      value: {name: 'value'},
      extraField1: {name: 'secure'},
    };
    component.dispatchersCallbacks = {addItem};

    component.addItem(0, '  param ', '', undefined, undefined, undefined, undefined, undefined, undefined, undefined);

    expect(addItem).toHaveBeenCalledOnceWith(12, 0, 'param', '');
  });
});
