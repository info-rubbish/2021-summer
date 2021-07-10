const config = {
    port: '3000',
    // when any byte of serviceworker.js file changed, changes will tigger onupdatefound event to renew the cache
    last_edit: '1625837814592', // result form Date.now()
}

self.addEventListener('install', (event) => {
    event.waitUntil(caches.delete('v1'))
})

self.addEventListener('activate', (event) => {
    console.log('serviceworker activated')
})

self.addEventListener('fetch', (event) => {
    const url = new URL(event.request.url)
    if (url.port != config.port)
        // if not under the port, send directly
        event.respondWith(Promise.resolve(fetch(event.request)))
    else if (url.hostname != 'localhost' && url.protocol == 'http:')
        // if browser unsupport, send directly
        event.respondWith(Promise.resolve(fetch(event.request)))
    else {
        event.respondWith(
            (async function () {
                var cache = await caches.open('v1')
                var response = await cache.match(event.request)
                if (!response) {
                    // console.log(`from Internet( ${event.request.url} )`)
                    var fetch_response = fetch(event.request)
                    var clone_response = (await fetch_response).clone()
                    if (!clone_response.ok)
                        throw new TypeError('Bad response status')
                    await cache.put(url, clone_response)
                    return fetch_response
                } else {
                    // console.log(`from Cache( ${event.request.url} )`)
                    return response
                }
            })()
        )
    }
})
