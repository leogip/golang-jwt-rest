import { updateAuthAlert } from './modules/alerts.js';
import { renderUserList } from './modules/render.js';
import { response, err } from './modules/response.js';

document.addEventListener('DOMContentLoaded', () => {
    updateAuthAlert();
    renderUserList(authByUID, false);

    function authByUID(e) {
        e.preventDefault();
        const url = e.target.getAttribute('href');

        fetch(url)
            .then((r) => r.json())
            .then((data) => {
                response(data);
                localStorage.setItem('token', data.data);
                updateAuthAlert();
            })
            .catch(err);
    }
});
