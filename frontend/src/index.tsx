import 'bootstrap/dist/css/bootstrap.css'
import './custom.scss'

import React from 'react'
import ReactDOM from 'react-dom'

import ApolloProvider from './apollo-provider'
import IntlProvider from './intl-provider'
import App from './app'

ReactDOM.render(
  <ApolloProvider>
    <IntlProvider>
      <App />
    </IntlProvider>
  </ApolloProvider>,
  document.getElementById('root'),
)
