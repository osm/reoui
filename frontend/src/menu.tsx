import React from 'react'
import { Collapse, Nav, Navbar, NavbarToggler, NavItem, NavLink } from 'reactstrap'
import { Link } from 'react-router-dom'
import { FormattedMessage } from 'react-intl'

const Menu: React.FC = () => {
  const [isOpen, setIsOpen] = React.useState<boolean>(false)

  return (
    <Navbar color="light" light expand="md">
      <NavbarToggler onClick={() => setIsOpen(!isOpen)} />
      <Collapse isOpen={isOpen} navbar>
        <Nav className="container-fluid" navbar>
          <NavItem>
            <NavLink tag={Link} to="/">
              <FormattedMessage id="menu.stream" defaultMessage="Stream" />
            </NavLink>
          </NavItem>
          <NavItem>
            <NavLink tag={Link} to="/videos">
              <FormattedMessage id="menu.videos" defaultMessage="Video archive" />
            </NavLink>
          </NavItem>
        </Nav>
      </Collapse>
    </Navbar>
  )
}

export default Menu
