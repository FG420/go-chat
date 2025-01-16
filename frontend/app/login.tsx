// LoginScreen.js
import { router } from 'expo-router';
import React, { useEffect, useState } from 'react';
import { View, Text, TextInput, Button, StyleSheet, Alert } from 'react-native';

const LoginScreen = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [valid, setValid] = useState(false)

    const handleLogin = () => {
        // Qui puoi aggiungere la logica per validare o fare login con un backend
        if (email === 'test@example.com' && password === 'password') {
            Alert.alert('Login Success', 'Benvenuto!');
            console.log('Login Success')
            setValid(true)
        } else {
            Alert.alert('Errore', 'Credenziali errate');
            console.log('Error')
            setValid(false)
        }
    };

    
    useEffect(() => {
        
        if (valid == true){
            router.navigate('/')
        }
    })

    return (
        <View style={styles.container}>
            <Text style={styles.title}>Accedi</Text>

            <TextInput
                style={styles.input}
                placeholder="Email"
                value={email}
                onChangeText={setEmail}
                keyboardType="email-address"
            />

            <TextInput
                style={styles.input}
                placeholder="Password"
                value={password}
                onChangeText={setPassword}
                secureTextEntry
            />

            <Button title="Login" onPress={handleLogin} />
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: 'center',
        padding: 20,
    },
    title: {
        fontSize: 24,
        fontWeight: 'bold',
        marginBottom: 20,
        textAlign: 'center',
    },
    input: {
        height: 40,
        borderColor: '#ccc',
        borderWidth: 1,
        marginBottom: 15,
        paddingLeft: 10,
        borderRadius: 5,
    },
});

export default LoginScreen;
