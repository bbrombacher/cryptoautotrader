import useSWR from 'swr'
import styles from '../styles/Home.module.css'

const fetcher = (...args) => fetch(...args).then((res) => res.json())

export default function TradeSessions() {
    const { data, error } = useSWR(
        ['https://cryptoautotrader-production.up.railway.app/v1/trade-sessions?limit=10', 
        {
        headers: {
           'x-user-id': 'd78964af-9dbe-4613-bf46-f3701bdd0494',
        }
    }], fetcher)

 
   if (error) return <p>could not fetch trade sessions</p>
   if (!data) return <p>no transaction data</p>

   return (
        <div  className={styles.tradesessions}>
            <h1> Your Trade Sessions </h1>
          <ul>
            {
                data.trade_sessions.map((trade_session) => 
                <li className={styles.tradesession} key={trade_session.id}>
                    <div> SessionID: {trade_session.id}</div>
                    <div> Starting Balance: {trade_session.starting_balance} </div>
                    <div> Ending Balance: {trade_session.ending_balance} </div>
                </li>
                )
            }
          </ul>
        </div>
    )

}

