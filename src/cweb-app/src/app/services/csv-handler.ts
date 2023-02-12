import {Injectable} from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class CsvService {
  public saveDataInCSV(data: Array<any>): string {
    if (data.length === 0) {
      return '';
    }

    const propertyNames = Object.keys(data[0]);
    let csvContent = propertyNames.join(',') + '\n';

    const rows: string[] = [];

    data.forEach((item) => {
      const values: string[] = [];

      propertyNames.forEach((key) => {
        let val: any = item[key];

        if (val !== undefined && val !== null) {
          val = String(val);
        } else {
          val = '';
        }
        values.push(val);
      });
      rows.push(values.join(','));
    });
    csvContent += rows.join('\n');

    return csvContent;
  }

  public importDataFromCSV(csvText: string): Array<any> {
    const propertyNames = csvText.slice(0, csvText.indexOf('\n')).split(/,|;/);
    const dataRows = csvText.slice(csvText.indexOf('\n') + 1).split('\n');

    const dataArray: any[] = [];
    dataRows.forEach((row) => {
      const values = row.split(/,|;/);
      if (values.length !== propertyNames.length) {
        return;
      }

      const obj: any = {};

      for (let index = 0; index < propertyNames.length; index++) {
        const propertyName: string = propertyNames[index].trim();

        let val: any = values[index].trim();
        if (val === '') {
          val = null;
        }

        obj[propertyName] = val;
      }
      dataArray.push(obj);
    });

    return dataArray;
  }

  public importDataFromCSVByType(csvText: string, obj: any): Array<any> {
    const propertyNames = csvText.slice(0, csvText.indexOf('\n')).split(',');
    const dataRows = csvText.slice(csvText.indexOf('\n') + 1).split('\n');


    const dataArray: any[] = [];
    dataRows.forEach((row) => {
      const values = row.split(',');

      const dataObj: any = {};
      for (let index = 0; index < propertyNames.length; index++) {
        const propertyName: string = propertyNames[index];

        let value: any = values[index];
        if (value === '') {
          value = null;
        }


        if (typeof obj[propertyName] === 'undefined') {
          dataObj[propertyName] = undefined;
        } else if (typeof obj[propertyName] === 'boolean') {
          dataObj[propertyName] = value.toLowerCase() === 'true';
        } else if (typeof obj[propertyName] === 'number') {
          dataObj[propertyName] = Number(value);
        } else if (typeof obj[propertyName] === 'string') {
          dataObj[propertyName] = value;
        } else if (typeof obj[propertyName] === 'object') {
          console.error('do no have algorithm to convert object');
        }
      }

      dataArray.push(dataObj);
    });

    return dataArray;
  }
}
