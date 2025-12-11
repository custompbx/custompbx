import {AfterViewInit, Component, ElementRef, EventEmitter, Input, OnInit, Output, ViewChild} from '@angular/core';
import {Ace, edit} from 'ace-builds';
// import 'ace-builds/webpack-resolver';
import 'ace-builds/src-noconflict/mode-xml';
// import 'ace-builds/src-noconflict/worker-xml';
import 'ace-builds/src-noconflict/theme-dracula';
import 'ace-builds/src-noconflict/ext-beautify';

import {MaterialModule} from "../../../material-module";

const THEME = 'ace/theme/dracula';
const LANG = 'ace/mode/xml';

@Component({
standalone: true,
    imports: [MaterialModule],
    selector: 'app-code-editor',
    templateUrl: './code-editor.component.html',
    styleUrls: ['./code-editor.component.css']
})
export class CodeEditorComponent implements OnInit, AfterViewInit {
  private codeEditor: Ace.Editor;
  private editorBeautify;
  @ViewChild('codeEditor') private codeEditorElmRef: ElementRef;
  @Input() content: string;
  @Output() contentChange = new EventEmitter();

  constructor() {
  }

  ngOnInit() {
  }
  ngAfterViewInit() {
    this.initEditor();
  }

  public initEditor() {
    const element = this.codeEditorElmRef.nativeElement;
    const editorOptions = this.getEditorOptions();
    this.codeEditor = this.createCodeEditor(element, editorOptions);
    this.setContent(this.content);
    // hold reference to beautify extension
    this.editorBeautify = require('ace-builds/src-noconflict/ext-beautify');
    this.MyOnContentChange();
  }


  private createCodeEditor(element: HTMLElement, options: any): Ace.Editor {
    const editor = edit(element, options);
    editor.setTheme(THEME);
    editor.getSession().setMode(LANG);
    editor.setShowFoldWidgets(true);
    return editor;
  }

  // missing propery on EditorOptions 'enableBasicAutocompletion' so this is a wolkaround still using ts
  private getEditorOptions(): Partial<Ace.EditorOptions> & { enableBasicAutocompletion?: boolean; } {
    const basicEditorOptions: Partial<Ace.EditorOptions> = {
      highlightActiveLine: true,
      minLines: 14,
      maxLines: Infinity,
      useWorker: false,
    };
    const extraEditorOptions = { enableBasicAutocompletion: true };
    return Object.assign(basicEditorOptions, extraEditorOptions);
  }

  /**
   * @returns - the current editor's content.
   */
  public getContent() {
    if (this.codeEditor) {
      const code = this.codeEditor.getValue();
      return code;
    }
  }

  /**
   * @param content - set as the editor's content.
   */
  public setContent(content: string): void {
    if (this.codeEditor) {
      if (!content) {
        content = '';
      }
      this.codeEditor.setValue(content);
    }
  }

  /**
   * @description
   *  beautify the editor content, rely on Ace Beautify extension.
   */
  public beautifyContent(): void {
    if (this.codeEditor && this.editorBeautify) {
      const session = this.codeEditor.getSession();
      this.editorBeautify.beautify(session);
    }
  }
  /**
   * @event OnContentChange - a proxy event to Ace 'change' event - adding additional data.
   * @param callback - recive the current content and 'change' event's original parameter.
   */
  public OnContentChange(callback: (content: string, delta: Ace.Delta) => void): void {
    this.codeEditor.on('change', (delta) => {
      const content = this.codeEditor.getValue();
      callback(content, delta);
    });
  }  /**
   * @event OnContentChange - a proxy event to Ace 'change' event - adding additional data.
   */
  public MyOnContentChange(): void {
    this.codeEditor.on('change', (delta) => {
      const content = this.codeEditor.getValue();
      this.contentChange.emit(content);
    });
  }
}
