// ChatScreen.tsx

import { Link } from "expo-router";
import { ScrollView, StyleSheet, Text, View } from "react-native";


const ChatScreen = () => {

    const fakeData = [
        {
            user: 'user1',
            lastMessage: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Officia, eos. Iste quae blanditiis illo provident fuga dolorem debitis. Eligendi ipsum cumque temporibus tempora accusantium iste, consequatur ipsam quasi iure optio!',
            date: new Date().toLocaleString()
        },
        {
            user: 'user2',
            lastMessage: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Officia, eos. Iste quae blanditiis illo provident fuga dolorem debitis. Eligendi ipsum cumque temporibus tempora accusantium iste, consequatur ipsam quasi iure optio!',
            date: new Date().toLocaleString()
        },
        {
            user: 'user3',
            lastMessage: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Officia, eos. Iste quae blanditiis illo provident fuga dolorem debitis. Eligendi ipsum cumque temporibus tempora accusantium iste, consequatur ipsam quasi iure optio!',
            date: new Date().toLocaleString()
        },
        {
            user: 'user4',
            lastMessage: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Officia, eos. Iste quae blanditiis illo provident fuga dolorem debitis. Eligendi ipsum cumque temporibus tempora accusantium iste, consequatur ipsam quasi iure optio!',
            date: new Date().toLocaleString()
        },
        {
            user: 'user5',
            lastMessage: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Officia, eos. Iste quae blanditiis illo provident fuga dolorem debitis. Eligendi ipsum cumque temporibus tempora accusantium iste, consequatur ipsam quasi iure optio!',
            date: new Date().toLocaleString()
        },
        {
            user: 'user6',
            lastMessage: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Officia, eos. Iste quae blanditiis illo provident fuga dolorem debitis. Eligendi ipsum cumque temporibus tempora accusantium iste, consequatur ipsam quasi iure optio!',
            date: new Date().toLocaleString()
        },
        {
            user: 'user7',
            lastMessage: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Officia, eos. Iste quae blanditiis illo provident fuga dolorem debitis. Eligendi ipsum cumque temporibus tempora accusantium iste, consequatur ipsam quasi iure optio!',
            date: new Date().toLocaleString()
        },
        {
            user: 'user8',
            lastMessage: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Officia, eos. Iste quae blanditiis illo provident fuga dolorem debitis. Eligendi ipsum cumque temporibus tempora accusantium iste, consequatur ipsam quasi iure optio!',
            date: new Date().toLocaleString()
        },
        {
            user: 'user9',
            lastMessage: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Officia, eos. Iste quae blanditiis illo provident fuga dolorem debitis. Eligendi ipsum cumque temporibus tempora accusantium iste, consequatur ipsam quasi iure optio!',
            date: new Date().toLocaleString()
        },
        {
            user: 'user10',
            lastMessage: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Officia, eos. Iste quae blanditiis illo provident fuga dolorem debitis. Eligendi ipsum cumque temporibus tempora accusantium iste, consequatur ipsam quasi iure optio!',
            date: new Date().toLocaleString()
        },
    ]



    return (
        <ScrollView >
            Chat Screeen
            <Link href={ '/chatter' } >Mona</Link>


            { fakeData.map( ( fake ) => (
                <View style={
                    styles.viewContainer
                }>
                    <Text style={ styles.user }>{ fake.user }</Text>
                    <Text style={ styles.msg } numberOfLines={ 1 } >{ fake.lastMessage }</Text>
                    <Text style={ styles.date }>{ fake.date }</Text>
                </View>
            ) ) }
        </ScrollView>
    )
}



const styles = StyleSheet.create( {
    viewContainer: {
        display: 'flex',
        justifyContent: 'center',
        alignContent: 'center',
        padding: 10
    },

    user: {
        display: 'flex',
        justifyContent: 'flex-start',
        alignItems: 'center',
        padding: 10
    },

    msg: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        padding: 10

    },

    date: {
        display: 'flex',
        justifyContent: 'flex-end',
        alignItems: 'center'
    },
} )



export default ChatScreen;

