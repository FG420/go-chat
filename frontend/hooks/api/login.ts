import axios from 'axios'


interface Login {
    username: string
    password: string
}


export const LoginCall = async ( data: Login ) => {
    try {

        const res = await axios.post( 'http://192.168.1.8:8080/login', data )
        // const res = await axios.get( "https://google.com", { headers: { "Access-Control-Allow-Origin": "*" } } )
        console.log( res )
        // const h = new Headers
        // h.set( "Content-Type", "application/json" )

        // const api = await fetch( "http://localhost:8080/login", { method: "POST", headers: h, body: JSON.stringify( data ) } )
        // console.log( api.body )
        // const res = await api.body()

    } catch ( err ) {
        console.log( err )
    }
}