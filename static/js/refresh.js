import { updateAuthAlert } from './modules/alerts.js';
import { authHeader } from './modules/authHeader.js';
import { response, err } from './modules/response.js';

document.addEventListener('DOMContentLoaded', () => {
    updateAuthAlert();
    const btn = document.querySelector('.btn-refresh');
    btn.addEventListener('click', refreshFetch);

    function refreshFetch(e) {
        e.preventDefault();
        const url = e.target.getAttribute('href');
        fetch(url, {
            method: 'GET',
            headers: authHeader(),
        })
            .then((r) => r.json())
            .then((data) => {
                response(data);
                localStorage.setItem('token', data.data || '');
                updateAuthAlert();
            })
            .catch(err);
    }
});
