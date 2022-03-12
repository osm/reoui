import React from 'react'
import moment from 'moment-timezone'
import { useQuery } from '@apollo/client'
import { Table, Button, Input, FormGroup, Label } from 'reactstrap'
import { FaPlay } from 'react-icons/fa'
import { Link } from 'react-router-dom'
import { useHistory, useParams } from 'react-router'
import { FormattedMessage } from 'react-intl'

import Download from './download'

import { ListVideos, ListVideosVariables } from './queries/__generated__/ListVideos'
import QUERY from './queries/ListVideos.graphql'

const Videos: React.FC = () => {
  const history = useHistory()
  const { date: paramDate }: { date?: string } = useParams()
  const [date, setDate] = React.useState(paramDate ?? moment().format('YYYY-MM-DD'))

  const { loading, data } = useQuery<ListVideos, ListVideosVariables>(QUERY, {
    variables: {
      date,
    },
    fetchPolicy: 'no-cache',
  })
  const videos = data?.videos ?? []

  return loading ? (
    <FormattedMessage id="videos.loading" defaultMessage="Loading videos..." />
  ) : (
    <>
      <FormGroup style={{ width: '250px' }}>
        <Label>
          <FormattedMessage id="videos.show-videos-for-date" defaultMessage="Show videos for date" />
        </Label>
        <Input
          type="date"
          value={date}
          onChange={(e) => {
            setDate(e.target.value)
            history.push(`/videos/${e.target.value}`)
          }}
        />
      </FormGroup>
      <Table style={{ marginTop: '10px' }}>
        <thead>
          <tr>
            <th>
              <FormattedMessage id="videos.camera-name" defaultMessage="Camera name" />
            </th>
            <th>
              <FormattedMessage id="videos.date" defaultMessage="Date" />
            </th>
            <th>
              <FormattedMessage id="videos.duration" defaultMessage="Duration" />
            </th>
            <th style={{ textAlign: 'right' }}>
              <FormattedMessage id="videos.actions" defaultMessage="Actions" />
            </th>
          </tr>
        </thead>
        <tbody>
          {videos.map((v) => {
            return (
              v && (
                <tr key={v.id}>
                  <td>{v.cameraName}</td>
                  <td>{moment(v.date).format('YYYY-MM-DD HH:mm:ss')}</td>
                  <td>
                    {v.duration} <FormattedMessage id="videos.seconds" defaultMessage="seconds" />
                  </td>
                  <td style={{ textAlign: 'right' }}>
                    <Button tag={Link} to={`/player/${v.id}`} color="primary" outline style={{ marginRight: '5px' }}>
                      <FaPlay /> <FormattedMessage id="videos.play" defaultMessage="Play" />
                    </Button>
                    <Download id={v.id} />
                  </td>
                </tr>
              )
            )
          })}
        </tbody>
      </Table>
    </>
  )
}

export default Videos
