

export function formatDate(date) {
    let dt = new Date(date)
    return dt.getDate() + '/' + dt.getMonth() + "/" + dt.getFullYear()
}