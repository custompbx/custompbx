import { NgModule, ModuleWithProviders } from '@angular/core';
import { CommonModule } from '@angular/common';

import { WebsocketService } from './websocket.service';
import { config } from './websocket.config';
import { WebSocketConfig } from './websocket.interfaces';


@NgModule({
    imports: [
        CommonModule
    ],
    declarations: [],
    providers: [
        WebsocketService
    ]
})
export class WebsocketModule {
    public static config(wsConfig: WebSocketConfig): ModuleWithProviders<WebsocketModule> {
        return {
            ngModule: WebsocketModule,
            providers: [{ provide: config, useValue: wsConfig }]
        };
    }
}
