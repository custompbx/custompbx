import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';

export interface IseqDiagram {
  [addr: string]: {
    startX: number,
    startY: number,
    endX: number,
    // endY: number,
    msgs: [{
      seq: number,
      to: string,
      date: string,
      name: string,
      body: string,
      name_cseq: string
    }],
  };
}

export interface IdataItem {
  hep_src_ip: string;
  hep_src_port: string;
  hep_dst_ip: string;
  hep_dst_port: string;
  hep_timestamp: string;
  sip_first_method: string;
  hep_payload: string;
  sip_cseq_method: string;
}

@Component({
  selector: 'app-svg-seq-diagram',
  templateUrl: './svg-seq-diagram.component.svg',
  styleUrls: ['./svg-seq-diagram.component.css']
})

export class SvgSeqDiagramComponent implements OnInit {

  public data: Array<IdataItem>;
  public startX = 80;
  public startY = 20;
  public endY = 300;
  public nodes: IseqDiagram = {};

  @Input()
  set inData(data: Array<IdataItem>) {
    this.data = data;
    this.structuring(this.data);
  }

  @Output() toView = new EventEmitter<string>();

  constructor() { }

  ngOnInit(): void {
  }

  structuring (data: Array<IdataItem>) {
    this.nodes = {};
    this.endY = data.length * 50 + 70;
    data.forEach(
      (msg, i) => {
        if (this.nodes[msg.hep_src_ip + ':' + msg.hep_src_port]) {
          this.nodes[msg.hep_src_ip + ':' + msg.hep_src_port].msgs.push({
            seq: i + 1,
            to: msg.hep_dst_ip + ':' + msg.hep_dst_port,
            date: msg.hep_timestamp,
            name: msg.sip_first_method,
            body: msg.hep_payload,
            name_cseq: msg.sip_cseq_method,
          });
          return;
        }
        this.nodes[msg.hep_src_ip + ':' + msg.hep_src_port] = {
          startX: this.getX(this.nodeCounter()),
          startY: this.startY,
          endX: this.getX(this.nodeCounter()),
          msgs: [{
            seq: i + 1,
            to: msg.hep_dst_ip + ':' + msg.hep_dst_port,
            date: msg.hep_timestamp,
            name: msg.sip_first_method,
            body: msg.hep_payload,
            name_cseq: msg.sip_cseq_method,
          }],
        };
      }
    );
  }

  getX (index: number): number {
    return this.startX + index * 170;
  }

  nodeCounter(): number {
    return Object.keys(this.nodes).length;
  }

  methodColor(method: string) {
    switch (method.charAt(0)) {
      case '1':
      case '2':
      case '3':
        return 'green';
      case '4':
      case '5':
      case '6':
        return 'red';
      default:
        return 'blue';
    }
  }

  methodName(msgs: any) {
    switch (msgs.name.charAt(0)) {
      case '1':
      case '2':
      case '3':
      case '4':
      case '5':
      case '6':
        return msgs.name + '(' + msgs.name_cseq + ')';
      default:
        return msgs.name;
    }
  }

  chooseToView(body) {
    this.toView.emit(body);
  }
}
