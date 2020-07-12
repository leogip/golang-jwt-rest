import { updateAuthAlert } from './modules/alerts.js';
import { renderUserList, renderTokenList } from './modules/render.js';
import { authHeader } from './modules/authHeader.js';
import { response, err } from './modules/response.js';

document.addEventListener('DOMContentLoaded', () => {
    updateAuthAlert();
    renderUserList(removeToken, true);
    renderTokenList(removeToken);

    function removeToken(e) {
        e.preventDefault();
        const url = e.target.getAttribute('href');

        fetch(url, {
            method: 'DELETE',
            headers: authHeader(),
        })
            .then((r) => r.json())
            .then((data) => {
                response(data);
                localStorage.removeItem('token');
                updateAuthAlert();
                if (e.target.href.match(/(\/t\/)/g)) e.target.remove();
                return;
            })
            .catch(err);
    }
});
