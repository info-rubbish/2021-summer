const config = {
    port: '3000',
    last_edit: '1625837814592', // result form Date.now()
}
self.addEventListener('install', (event) => {
    event.waitUntil(
        //maybe promise chain is better
        // clear cache
        caches.delete('v1').then(() => {
            caches.open('v1').then((cache) => {
                // Add pre-cache path below
                return cache.addAll([])
            })
        })
    )
})
self.addEventListener('activate', (event) => {
    console.log('serviceworker activated')
})

self.addEventListener('fetch', (event) => {
    const url = new URL(event.request.url)
    // console.log(url.host)
    if (url.port != config.port)
        // if not under the port, send directly
        event.respondWith(Promise.resolve(fetch(event.request)))
    else if (url.hostname != 'localhost' && url.protocol == 'http:')
        // if browser unsupport, send directly
        event.respondWith(Promise.resolve(fetch(event.request)))
    else {
        //if under port cache request or send existed cache
        event.respondWith(
            caches.match(event.request).then((response) => {
                if (!response) {
                    console.log(`from Internet( ${event.request.url} )`)
                    caches.open('v1').then((cache) => {
                        fetch(url).then(function (response) {
                            if (!response.ok)
                                throw new TypeError('Bad response status')
                            return cache.put(url, response)
                        })
                    })
                } else {
                    console.log(`from cache( ${event.request.url} )`)
                    return response
                }
            })
        )
    }
})
