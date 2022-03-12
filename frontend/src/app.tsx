import React from 'react'
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'

import NotFound from './not-found'
import Videos from './videos'
import Stream from './stream'
import Player from './player'
import Menu from './menu'

const App: React.FC = () => {
  return (
    <div className="container-fluid">
      <Router>
        <Menu />
        <div className="container mt-4">
          <Switch>
            <Route exact path="/player/:id">
              <Player />
            </Route>
            <Route exact path="/">
              <Stream />
            </Route>
            <Route exact path="/stream/:id">
              <Stream />
            </Route>
            <Route path="/videos/:date?">
              <Videos />
            </Route>
            <Route path="*">
              <NotFound />
            </Route>
          </Switch>
        </div>
      </Router>
    </div>
  )
}

export default App
