import { updateAuthAlert, showAlert } from './alerts.js';

export function response(data) {
    if (!data.success) {
        showAlert('warning', data.msg);
        return;
    }
    showAlert('success', data.msg);
}

export function err() {
    showAlert('warning', 'Server error');
    localStorage.removeItem('token');
    updateAuthAlert();
    return;
}
