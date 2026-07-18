import {Component, inject, linkedSignal, signal} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {Store} from '@ngrx/store';
import {TranslocoPipe} from '@jsverse/transloco';
import {AppState} from '../../store/app.states';
import {
  ClearWebUserAvatar,
  UpdateWebUserAvatar,
  UpdateWebUserPassword,
  UpdateWebUserStun,
  UpdateWebUserVertoWs,
  UpdateWebUserWebRTCLib,
  UpdateWebUserWs,
} from '../../store/settings/settings.actions';
import {UserService} from '../../services/user.service';
import {LocaleService} from '../../i18n/locale.service';
import {InnerHeaderComponent} from '../inner-header/inner-header.component';
import {IconComponent} from '../icon/icon.component';

@Component({
  standalone: true,
  selector: 'app-profile',
  imports: [FormsModule, InnerHeaderComponent, IconComponent, TranslocoPipe],
  templateUrl: './profile.component.html',
  styleUrl: './profile.component.css',
})
export class ProfileComponent {
  private readonly store = inject(Store<AppState>);
  readonly user = inject(UserService).userSignal;
  readonly localeService = inject(LocaleService);
  readonly acceptFile = 'image/png,image/jpeg,image/webp';
  readonly password = signal('');
  readonly ws = linkedSignal(() => this.user()?.ws ?? '');
  readonly vertoWs = linkedSignal(() => this.user()?.verto_ws ?? '');
  readonly stun = linkedSignal(() => this.user()?.stun ?? '');
  readonly technicalLabels = {
    webRtc: 'WebRTC',
    sipJs: 'Sip.js',
    verto: 'Verto',
  } as const;

  changeLocale(value: string): void {
    this.localeService.setLocale(value);
  }

  chooseAvatar(input: HTMLInputElement): void {
    const file = input.files?.[0];
    const user = this.user();
    if (!file || !user?.id || file.size > 512000) {
      input.value = '';
      return;
    }
    const reader = new FileReader();
    reader.onload = () => this.store.dispatch(new UpdateWebUserAvatar({file: reader.result, id: user.id}));
    reader.readAsDataURL(file);
    input.value = '';
  }

  clearAvatar(): void {
    const user = this.user();
    if (user?.id) {
      this.store.dispatch(new ClearWebUserAvatar({file: '', id: user.id}));
    }
  }

  updatePassword(): void {
    const user = this.user();
    const password = this.password();
    if (user?.id && password.length >= 6) {
      this.store.dispatch(new UpdateWebUserPassword({id: user.id, password}));
      this.password.set('');
    }
  }

  updateWebRtcLibrary(value: string): void {
    const user = this.user();
    if (user?.id && value !== user.webrtc_lib) {
      this.store.dispatch(new UpdateWebUserWebRTCLib({id: user.id, value}));
    }
  }

  updateWs(): void {
    const user = this.user();
    if (user?.id && this.ws() !== (user.ws ?? '')) {
      this.store.dispatch(new UpdateWebUserWs({id: user.id, value: this.ws()}));
    }
  }

  updateVertoWs(): void {
    const user = this.user();
    if (user?.id && this.vertoWs() !== (user.verto_ws ?? '')) {
      this.store.dispatch(new UpdateWebUserVertoWs({id: user.id, value: this.vertoWs()}));
    }
  }

  updateStun(): void {
    const user = this.user();
    if (user?.id && this.stun() !== (user.stun ?? '')) {
      this.store.dispatch(new UpdateWebUserStun({id: user.id, value: this.stun()}));
    }
  }
}
