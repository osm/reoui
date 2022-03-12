import React from 'react'
import { ApolloClient, InMemoryCache } from '@apollo/client'
import { createUploadLink } from 'apollo-upload-client'
import { ApolloProvider as ReactApolloProvider } from '@apollo/client/react'

const graphqlUrl = process.env.GRAPHQL_URL || '/graphql'

const httpLink = createUploadLink({
  uri: graphqlUrl,
})

const client = new ApolloClient({
  link: httpLink,
  cache: new InMemoryCache(),
})

const ApolloProvider: React.FC = ({ children }: { children?: React.ReactNode }) => {
  return (
    <>
      <ReactApolloProvider client={client}>{children}</ReactApolloProvider>
    </>
  )
}

export default ApolloProvider
