const $elResponseAlert = document.querySelector('.response');
const $elAuthAlert = document.querySelector('.authorization');

export function updateAuthAlert() {
    if (localStorage.getItem('token')) {
        $elAuthAlert.style.display = 'block';
        $elAuthAlert.querySelector(
            '.msg'
        ).textContent = `Your token: ${localStorage.getItem('token')}`;
    } else {
        $elAuthAlert.style.display = 'none';
    }
}

export function showAlert(type, msg) {
    $elResponseAlert.style.display = 'block';
    $elResponseAlert.classList.add(`alert-${type}`);
    $elResponseAlert.querySelector('.alert-heading').textContent = type;
    $elResponseAlert.querySelector('.msg').textContent = msg;
}
