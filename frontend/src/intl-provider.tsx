import React from 'react'
import { IntlProvider as ReactIntlProvider } from 'react-intl'

const IntlProvider: React.FC = ({ children }: { children?: React.ReactNode }) => {
  const urlSearchParams = new URLSearchParams(window.location.search)
  const params = Object.fromEntries(urlSearchParams.entries())
  const { locale: queryLocale } = params

  const requestedLocale = queryLocale || localStorage.getItem('__LOCALE__') || 'en'
  const locale = requestedLocale in messages ? requestedLocale : 'en'

  if (localStorage.getItem('__LOCALE__') !== locale) {
    localStorage.setItem('__LOCALE__', locale)
  }

  return (
    <>
      <ReactIntlProvider messages={messages[locale]} locale={locale} defaultLocale="en">
        {children}
      </ReactIntlProvider>
    </>
  )
}

export default IntlProvider

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const messages: any = {
  sv: {
    'menu.stream': 'Direktsändning',
    'menu.videos': 'Filmarkiv',

    'videos.loading': 'Laddar filmer...',
    'videos.show-videos-for-date': 'Visa filmer för datum',
    'videos.camera-name': 'Kamera',
    'videos.date': 'Datum',
    'videos.duration': 'Längd',
    'videos.actions': 'Hantera',
    'videos.seconds': 'sekunder',
    'videos.play': 'Spela',

    'download.download': 'Ladda hem',

    'player.pause': 'Pausa',
    'player.play': 'Spela',

    'stream.not-configured': 'Kameran är inte konfigurerad',
    'stream.select-camera': 'Välj kamera',
  },
}
