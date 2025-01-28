import { DarkTheme, DefaultTheme, ThemeProvider } from '@react-navigation/native';
import { Stack } from 'expo-router';
import { StatusBar } from 'expo-status-bar';
import React from 'react';
import { useColorScheme } from 'react-native';

export default function RootLayout () {
    const colorScheme = useColorScheme()
    return (
        <ThemeProvider value={ colorScheme === 'dark' ? DarkTheme : DefaultTheme }>
            <Stack screenOptions={ {
                headerShown: false, contentStyle: {
                    backgroundColor: 'transparent'
                }
            } }>
                <Stack.Screen name='(tabs)' options={ {
                } } />
                {/* <Stack.Screen name="chat"/> */ }
                {/* <Stack.Screen name="index" /> */ }
                {/* <Stack.Screen name="login" /> */ }
            </Stack>
            <StatusBar style='auto' />

        </ThemeProvider>
    );
}
