import React, { useState } from 'react';
import { View, Text, TextInput, Button, FlatList, StyleSheet, KeyboardAvoidingView, Platform } from 'react-native';

// Definiamo il tipo del messaggio
interface Message {
    id: string;
    text: string;
    sender: 'me' | 'other'; // Differenzia il mittente del messaggio
}

const ChatComponent = () => {
    const [ message, setMessage ] = useState<string>( '' );
    const [ messages, setMessages ] = useState<Message[]>( [
        { id: '1', text: 'Ciao! Come posso aiutarti oggi?', sender: 'other' },
        { id: '2', text: 'Salve, vorrei maggiori informazioni.', sender: 'me' },
    ] );
    const [ users, setUsers ] = useState<string[]>( [ 'User 1', 'User 2', 'User 3' ] ); // Lista utenti connessi


    const appendMsg = ( text: string ) => {

    }

    if ( window[ "WebSocket" ] ) {
        var conn = new WebSocket( "ws://" + window.location.host + "/ws" )
        conn.onclose = function ( event ) {

        }
    }

    // Funzione per inviare il messaggio
    const handleSendMessage = () => {
        if ( message.trim() === '' ) return; // Evita inviare messaggi vuoti

        const newMessage: Message = {
            id: ( messages.length + 1 ).toString(),
            text: message,
            sender: 'me', // Il messaggio è inviato dall'utente corrente
        };

        setMessages( ( prevMessages ) => [ ...prevMessages, newMessage ] );
        setMessage( '' ); // Resetta il campo input
    };

    // Funzione per simulare un messaggio ricevuto da un altro utente
    const handleReceiveMessage = () => {
        const newMessage: Message = {
            id: ( messages.length + 1 ).toString(),
            text: 'Ciao, come va?',
            sender: 'other',
        };

        setMessages( ( prevMessages ) => [ ...prevMessages, newMessage ] );
    };

    return (
        <KeyboardAvoidingView
            style={ styles.container }
            behavior={ Platform.OS === 'ios' ? 'padding' : 'height' }
        >
            <View style={ styles.chatContainer }>
                {/* Sezione per i messaggi */ }
                <FlatList
                    data={ messages }
                    renderItem={ ( { item } ) => (
                        <View
                            style={ [
                                styles.messageContainer,
                                {
                                    backgroundColor: item.sender === 'me' ? '#003366' : '#ccc', // Colore diverso per me e gli altri
                                    alignSelf: item.sender === 'me' ? 'flex-end' : 'flex-start', // Posizione diversa per me e gli altri
                                },
                            ] }
                        >
                            <Text style={ styles.messageText }>{ item.text }</Text>
                        </View>
                    ) }
                    keyExtractor={ ( item ) => item.id }
                    style={ styles.messagesList }

                />

                {/* Sezione per l'invio dei messaggi */ }
                <View style={ styles.inputContainer }>
                    <TextInput
                        style={ styles.input }
                        placeholder="Scrivi un messaggio..."
                        value={ message }
                        onChangeText={ setMessage }
                    />
                    <Button title="Invia" onPress={ handleSendMessage } />
                </View>
            </View>

            {/* Sezione per gli utenti connessi */ }
            <View style={ styles.usersContainer }>
                <Text style={ styles.usersTitle }>Utenti Connessi</Text>
                <FlatList
                    data={ users }
                    renderItem={ ( { item } ) => <Text style={ styles.userItem }>{ item }</Text> }
                    keyExtractor={ ( item, index ) => index.toString() }
                />
                <Button title="Simula messaggio ricevuto" onPress={ handleReceiveMessage } />
            </View>
        </KeyboardAvoidingView>
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
