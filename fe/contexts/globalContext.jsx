import React, { useState } from 'react';

export const GlobalContext = React.createContext({
    userID: "",
    tradeSessionID: "",
    setUserID: () => { console.log('executing set user id') },
    setTradeSessionID: () => { console.log('executing set user id') },
})

export const GlobalContextProvider = (props) => {
    const [userID, setUserID] = useState({});
    const [tradeSessionID, setTradeSessionID] = useState({})

    return (
        <GlobalContext.Provider 
            value={{
                userID: userID,
                setUserID: setUserID,
                tradeSessionID: tradeSessionID,
                setTradeSessionID: setTradeSessionID,
            }}
        >
            {props.children}
        </GlobalContext.Provider>
    )
}