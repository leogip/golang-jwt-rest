export async function getList(target) {
    let response = await fetch(`/api/${target}`);
    let data = await response.json();
    return data;
}
