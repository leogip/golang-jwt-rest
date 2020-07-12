import { getList } from './getList.js';

function createItem(url, label) {
    const item = document.createElement('a');
    item.classList.add('u-item');
    item.classList.add('list-group-item');
    item.classList.add('list-group-item-action');
    item.href = url;
    item.textContent = label;
    return item;
}

export function renderUserList(fnEvent, remove) {
    getList('users').then((ulist) => {
        const container = document.querySelector('.u-list');
        ulist.data.map((u) => {
            const url = remove
                ? `/api/token/remove/u/${u._id}`
                : `/api/token/get/${u._id}`;
            const label = `${u.fullname} id: ${u._id}`;
            const node = createItem(url, label);
            node.onclick = fnEvent;
            container.append(node);
        });
    });
}

export function renderTokenList(fnEvent) {
    getList('tokens').then((tlist) => {
        const container = document.querySelector('.t-list');
        tlist.data.map((t) => {
            const url = `/api/token/remove/t/${t._id}`;
            const label = `id: ${t._id}`;
            const node = createItem(url, label);
            node.onclick = fnEvent;
            container.append(node);
        });
    });
}
