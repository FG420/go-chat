


interface Login {
    username: string
    password: string
}


export const LoginCall = async ( data: Login ) => {
    try {
        const h = new Headers
        h.set( "Content-Type", "application/json" )

        const api = await fetch( "http://localhost:8080/login", { method: "POST", headers: h, body: JSON.stringify( data ) } )
        console.log( api.body )
        // const res = await api.body()

    } catch ( err ) {
        console.log( err )
    }
}