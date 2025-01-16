import { Stack } from 'expo-router';
import { StatusBar } from 'expo-status-bar';
import React from 'react';

export default function RootLayout() {
  return (
    <>
      <Stack screenOptions={{ headerShown: false }}>
        <Stack.Screen name='(tabs)' />
        {/* <Stack.Screen name="chat"/> */}
        {/* <Stack.Screen name="index" /> */}
        {/* <Stack.Screen name="login" /> */}
      </Stack>
      <StatusBar style='auto' />

    </>
  );
}
