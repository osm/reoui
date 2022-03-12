import React from 'react'
import { Button } from 'reactstrap'
import { FaRegSave } from 'react-icons/fa'
import { FormattedMessage } from 'react-intl'

const download = (url: string, filename: string): Promise<void> =>
  fetch(url).then((res) =>
    res.blob().then((blob) => {
      const url = window.URL.createObjectURL(blob)
      const link = document.createElement('a')
      link.href = url
      link.setAttribute('download', filename)
      document.body.appendChild(link)
      link.click()
      link?.parentNode?.removeChild(link)
    }),
  )

const Download: React.FC<{
  id: string
}> = ({ id }) => {
  return (
    <Button color="primary" outline onClick={() => download(`/files/${id}`, id)}>
      <FaRegSave /> <FormattedMessage id="download.download" defaultMessage="Download" />
    </Button>
  )
}

export default Download
