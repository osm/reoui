import React from 'react'
import { ReactFlvPlayer } from 'react-flv-player'
import { useQuery } from '@apollo/client'
import { Dropdown, DropdownToggle, DropdownMenu, DropdownItem } from 'reactstrap'
import { useHistory, useParams } from 'react-router'
import { FormattedMessage } from 'react-intl'

import { ListCameras } from './queries/__generated__/ListCameras'
import QUERY from './queries/ListCameras.graphql'

const Stream: React.FC = () => {
  const history = useHistory()
  const { id = '0' }: { id: string } = useParams()

  const { data } = useQuery<ListCameras>(QUERY, {
    fetchPolicy: 'no-cache',
  })
  const cameras = data?.cameras ?? []

  const [isOpen, setIsOpen] = React.useState(false)
  const [cameraId, setCameraId] = React.useState(id)

  return !cameras.find((c) => c?.id === id) ? (
    <FormattedMessage id="stream.not-configured" defaultMessage="This camera is not configured" />
  ) : (
    <>
      {cameras.length > 1 && (
        <Dropdown style={{ marginBottom: '5px' }} isOpen={isOpen} toggle={() => setIsOpen(!isOpen)}>
          <DropdownToggle caret>
            <FormattedMessage id="stream.select-camera" defaultMessage="Select camera" />
          </DropdownToggle>
          <DropdownMenu>
            {cameras?.map(
              (c) =>
                c && (
                  <DropdownItem
                    size="sm"
                    key={c.id}
                    onClick={() => {
                      history.push(`/stream/${c.id}`)
                      setCameraId(c.id)
                    }}
                  >
                    {c.name}
                  </DropdownItem>
                ),
            )}
          </DropdownMenu>
        </Dropdown>
      )}
      <ReactFlvPlayer
        key={`camera-${cameraId}`}
        hasAudio={true}
        hasVideo={true}
        isMuted={false}
        url={`/stream/${cameraId}`}
      />
    </>
  )
}

export default Stream
