import React from 'react'
import { Container, Layout, Icon, Color, Text, FontVariation } from '@harness/uicore'
import { Link } from 'react-router-dom'
import { useStrings } from 'framework/strings'
import { useAppContext } from 'AppContext'
import { useGetSpaceParam } from 'hooks/useGetSpaceParam'

import type { GitInfoProps } from 'utils/GitUtils'
import css from './RepositorySettingsHeader.module.scss'

export function RepositorySettingsHeader({ repoMetadata }: Pick<GitInfoProps, 'repoMetadata'>) {
  const { getString } = useStrings()
  const { routes } = useAppContext()
  const space = useGetSpaceParam()
  return (
    <Container className={css.header}>
      <Container>
        <Layout.Horizontal spacing="small" className={css.breadcrumb}>
          <Link to={routes.toSCMRepositoriesListing({ space })}>{getString('repositories')}</Link>
          <Icon name="main-chevron-right" size={10} color={Color.GREY_500} />
          <Link to={routes.toSCMRepository({ repoPath: repoMetadata.path as string })}>{repoMetadata.uid}</Link>
        </Layout.Horizontal>
        <Container padding={{ top: 'medium', bottom: 'medium' }}>
          <Text font={{ variation: FontVariation.H4 }}>{getString('settings')}</Text>
        </Container>
      </Container>
    </Container>
  )
}
