import React, { useState, useEffect, useRef } from 'react';
import { View, Text, TextInput, Button, ScrollView, StyleSheet, NativeSyntheticEvent, TextInputChangeEventData } from 'react-native';

const App: React.FC = () => {
    const [ msg, setMsg ] = useState<string>( '' )
    const [ messages, setMessages ] = useState<iMessage[]>( [] )
    const logRef = useRef<ScrollView>( null )
    const conn = useRef<WebSocket | null>( null )
    const [ users, setUsers ] = useState<string[]>( [] )

    useEffect( () => {
        if ( window.WebSocket ) {
            conn.current = new WebSocket( 'ws://localhost:8080/ws' );

            conn.current.onmessage = ( evt: MessageEvent ) => {
                const receivedMessage = evt.data.trim();
                console.log( JSON.parse( evt.data ) );

                const message: iMessage = JSON.parse( evt.data )

                if ( receivedMessage.includes( "Connected" ) ) {
                    setMessages( ( prevMessages ) => [
                        ...prevMessages,
                        { id: prevMessages.length, text: message.text, sender: 'System', timestamp: "", user: message.user }, // Etichetta come System
                    ] );

                    setUsers( ( prevUsers ) => [
                        ...prevUsers, message.user!
                    ] )
                    console.log( messages )
                    console.log( users )

                } else {
                    // Messaggio normale

                    console.log( "normale", messages )
                    setMessages( ( prevMessages ) => [
                        ...prevMessages,
                        { id: prevMessages.length, text: message.text, sender: 'Other', timestamp: message.timestamp, user: message.user },
                    ] );

                    setUsers( ( prevUsers ) => [
                        ...prevUsers, message.user!
                    ] )

                    console.log( messages )
                    console.log( users )
                }
            };

            conn.current.onclose = () => {
                setMessages( ( prevMessages ) => [
                    ...prevMessages,
                    { id: prevMessages.length, text: 'Connection closed.', sender: 'System', timestamp: "", user: null },
                ] );
            };
        } else {
            setMessages( ( prevMessages ) => [
                ...prevMessages,
                { id: prevMessages.length, text: 'Your browser does not support WebSockets.', sender: 'System', timestamp: "", user: null },
            ] );
        }

        return () => {
            if ( conn.current ) {
                conn.current.close();
            }
        };
    }, [] );


    // Funzione per inviare il messaggio
    const handleSubmit = () => {
        if ( conn.current && msg.trim() ) {

            const mex: iMessage = {
                user: messages[ 0 ].user,
                text: msg.trim(),
                timestamp: Date.now(),
                id: messages[ 0 ].id,
                sender: 'You'
            }
            console.log( "handle", JSON.stringify( mex ) )
            // const mex: iMessage = {
            //     user : 
            // }

            conn.current.send( JSON.stringify( mex ) );

            setMessages( ( prevMessages ) => [
                ...prevMessages,
                { id: prevMessages.length, text: msg, sender: 'You', timestamp: new Date().toLocaleString(), user: null },
            ] );
            setMsg( '' );
        }
    };

    const handleInputChange = ( e: NativeSyntheticEvent<TextInputChangeEventData> ) => {
        setMsg( e.nativeEvent.text );
    };

    return (
        <View style={ styles.container }>
            <View style={ styles.chatContainer }>
                <ScrollView
                    style={ styles.messagesList }
                    ref={ logRef }
                    onContentSizeChange={ () => {
                        logRef.current?.scrollToEnd( { animated: true } );
                    } }
                >
                    {/* Visualizzare i messaggi */ }
                    { messages.map( ( message ) => (
                        <View
                            key={ message.id }
                            style={ [
                                styles.messageContainer,
                                message.sender === 'You' ? styles.myMessage : message.sender === 'System' ? styles.systemMessage : styles.otherMessage,
                            ] }
                        >
                            { message.sender === "System" ? ( null ) : (
                                <Text style={ styles.messageSender }>{ message.sender }</Text>
                            ) }
                            <Text style={ styles.messageText }>{ message.text }</Text>
                            { message.timestamp === "" ? ( null ) : (
                                <Text style={ styles.messageText }>{ new Date( message.timestamp ).toLocaleTimeString() }</Text>
                            ) }

                        </View>
                    ) ) }
                </ScrollView>

                <View style={ styles.inputContainer }>
                    <TextInput
                        style={ styles.input }
                        value={ msg }
                        onChange={ handleInputChange }
                        placeholder="Type a message"
                        autoFocus
                    />
                    <Button title="Send" onPress={ handleSubmit } />
                </View>
            </View>

            <View style={ styles.usersContainer }>
                <Text style={ styles.usersTitle }>Users</Text>
                {/* Aggiungere la lista degli utenti (puoi implementare il meccanismo sul server) */ }
                { messages.map( ( message ) => (
                    <Text>{ message.sender }: { message.user }</Text>
                ) ) }

            </View>
        </View>
    );
};

const styles = StyleSheet.create( {
    container: {
        flex: 1,
        flexDirection: 'row',
        backgroundColor: '#f0f0f0',
    },
    chatContainer: {
        flex: 3,
        justifyContent: 'space-between',
        paddingTop: 10,
        backgroundColor: '#fff',
    },
    messagesList: {
        flex: 1,
        paddingHorizontal: 10,
        paddingBottom: 20,
    },
    messageContainer: {
        padding: 10,
        borderRadius: 10,
        marginVertical: 5,
        maxWidth: '80%',
    },
    myMessage: {
        backgroundColor: '#4CAF50', // Verde per i messaggi inviati dall'utente
        alignSelf: 'flex-end',
    },
    otherMessage: {
        backgroundColor: '#1976D2', // Blu più scuro per i messaggi ricevuti
        alignSelf: 'flex-start',
    },
    systemMessage: {
        backgroundColor: '#D3D3D3', // Grigio chiaro per i messaggi di sistema
        alignSelf: 'center',
        color: '#000',
        fontStyle: 'italic',
        maxWidth: '90%',
        paddingVertical: 10,
        paddingHorizontal: 15,
        borderRadius: 15,
    },
    messageSender: {
        fontWeight: 'bold',
        fontSize: 14,
        marginBottom: 5,
        color: '#333',
    },
    messageText: {
        color: 'white',
        fontSize: 16,
    },
    inputContainer: {
        flexDirection: 'row',
        padding: 10,
        alignItems: 'center',
        backgroundColor: '#fff',
        borderTopWidth: 1,
        borderTopColor: '#ccc',
    },
    input: {
        flex: 1,
        height: 40,
        borderColor: '#ccc',
        borderWidth: 1,
        borderRadius: 5,
        paddingLeft: 10,
        marginRight: 10,
    },
    usersContainer: {
        flex: 1,
        backgroundColor: '#fff',
        padding: 10,
        borderLeftWidth: 1,
        borderLeftColor: '#ccc',
    },
    usersTitle: {
        fontSize: 18,
        fontWeight: 'bold',
        marginBottom: 10,
    },
    userItem: {
        fontSize: 16,
        marginBottom: 5,
    },
} );

export default App;
