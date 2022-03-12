import React from 'react'
import { Player as VideoPlayer } from 'video-react'
import { Button } from 'reactstrap'
import { useParams } from 'react-router'
import { FaPlay, FaPause } from 'react-icons/fa'
import { FormattedMessage } from 'react-intl'

import Download from './download'

const Player: React.FC = () => {
  const { id }: { id: string } = useParams()
  const ref = React.useRef()
  const [isPlaying, setIsPlaying] = React.useState(false)

  return (
    <div style={{ marginBottom: '10px' }}>
      <div style={{ marginBottom: '10px' }}>
        <Button
          outline
          color="primary"
          style={{ marginRight: '5px' }}
          onClick={() => {
            setIsPlaying(!isPlaying)
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            // @ts-ignore
            ref?.current?.actions.togglePlay()
          }}
        >
          {isPlaying ? (
            <span>
              <FaPause /> <FormattedMessage id="player.pause" defaultMessage="Pause" />
            </span>
          ) : (
            <span>
              <FaPlay /> <FormattedMessage id="player.play" defaultMessage="Play" />
            </span>
          )}
        </Button>
        <Download id={id} />
      </div>
      <VideoPlayer playsInline src={`/files/${id}`} ref={ref} />
    </div>
  )
}

export default Player
