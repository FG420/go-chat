import { Tabs } from 'expo-router';
import React from 'react';
import Ionicon from '@expo/vector-icons/Ionicons'
import MCIcon from '@expo/vector-icons/MaterialCommunityIcons'


export default function TabLayout () {

    return (
        <Tabs
            screenOptions={ {
                headerShown: false,
                tabBarLabelPosition: 'below-icon',
                tabBarStyle: {
                    height: 'auto',
                    borderRadius: 24,
                    marginHorizontal: 20
                },
                // tabBarShowLabel: false
                animation: 'shift',
                tabBarBadgeStyle: {
                    color: 'white',
                    backgroundColor: 'blue'
                }
            } }>
            <Tabs.Screen
                name="chat"
                options={ {
                    title: 'Chat',
                    // headerShown: true,
                    // headerSearchBarOptions: {
                    //     placeholder: "Search the chat of choice"
                    // },
                    tabBarIcon: ( { color, focused } ) => <MCIcon size={ 28 } name={ focused ? 'chat-alert' : 'chat-outline' } color={ color } />,
                } }
            />
            <Tabs.Screen
                name="group"
                options={ {
                    title: 'Group Chat',
                    tabBarIcon: ( { color, focused } ) => <MCIcon size={ 28 } name={ focused ? 'account-group' : 'account-group-outline' } color={ color } />,
                } }
            />
            <Tabs.Screen
                name="index"
                options={ {
                    title: 'Home',
                    tabBarIcon: ( { color, focused } ) => <Ionicon size={ 28 } name={ focused ? 'home-sharp' : 'home-outline' } color={ color } />,
                } }
            />
            <Tabs.Screen
                name="search"
                options={ {
                    title: 'Search',
                    tabBarIcon: ( { color, focused } ) => <Ionicon size={ 28 } name={ focused ? 'search-sharp' : 'search-outline' } color={ color } />,
                } }
            />
            {/* <Tabs.Screen
                name="profile"
                options={ {
                    title: 'Profile',
                    tabBarIcon: ( { color, focused } ) => <FAIcon size={ 28 } name={ focused ? 'user-circle' : 'user-circle-o' } color={ color } />,
                } }
            /> */}
        </Tabs>
    );
}