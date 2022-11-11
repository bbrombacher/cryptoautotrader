import React, { useState } from 'react';

export const GlobalContext = React.createContext({
    userID: "",
    setUserID: () => { console.log('executing set user id') },
})

export const GlobalContextProvider = (props) => {
    const [userID, setUserID] = useState({});

    return (
        <GlobalContext.Provider 
            value={{
                userID: userID,
                setUserID: setUserID,
            }}
        >
            {props.children}
        </GlobalContext.Provider>
    )
}