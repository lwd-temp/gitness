import React, { useCallback, useState } from 'react'
import {
  Container,
  Text,
  FontVariation,
  Layout,
  Button,
  ButtonVariation,
  ButtonSize,
  TextInput,
  useToaster
} from '@harness/uicore'
import { useMutate } from 'restful-react'
import { Match, Truthy, Else } from 'react-jsx-match'
import { useStrings } from 'framework/strings'
import { ButtonRoleProps, getErrorMessage } from 'utils/Utils'
import type { GitInfoProps } from 'utils/GitUtils'
import type { TypesPullReq } from 'services/code'
import { PipeSeparator } from 'components/PipeSeparator/PipeSeparator'
import css from './PullRequest.module.scss'

interface PullRequestTitleProps extends TypesPullReq, Pick<GitInfoProps, 'repoMetadata'> {
  onSaveDone?: (newTitle: string) => Promise<boolean>
  onAddDescriptionClick: () => void
}

export const PullRequestTitle: React.FC<PullRequestTitleProps> = ({
  repoMetadata,
  title,
  number,
  description,
  onAddDescriptionClick
}) => {
  const [original, setOriginal] = useState(title)
  const [val, setVal] = useState(title)
  const [edit, setEdit] = useState(false)
  const { getString } = useStrings()
  const { showError } = useToaster()
  const { mutate } = useMutate({
    verb: 'PATCH',
    path: `/api/v1/repos/${repoMetadata.path}/+/pullreq/${number}`
  })
  const submitChange = useCallback(() => {
    mutate({
      title: val,
      description
    })
      .then(() => {
        setEdit(false)
        setOriginal(val)
      })
      .catch(exception => showError(getErrorMessage(exception), 0))
  }, [description, val, mutate, showError])

  return (
    <Layout.Horizontal spacing="small" className={css.prTitle}>
      <Match expr={edit}>
        <Truthy>
          <Container>
            <Layout.Horizontal spacing="small">
              <TextInput
                wrapperClassName={css.input}
                value={val}
                onFocus={event => event.target.select()}
                onInput={event => setVal(event.currentTarget.value)}
                autoFocus
                onKeyDown={event => {
                  switch (event.key) {
                    case 'Enter':
                      submitChange()
                      break
                    case 'Escape': // does not work, maybe TextInput cancels ESC?
                      setEdit(false)
                      break
                  }
                }}
              />
              <Button
                variation={ButtonVariation.PRIMARY}
                text={getString('save')}
                size={ButtonSize.MEDIUM}
                disabled={(val || '').trim().length === 0 || title === val}
                onClick={submitChange}
              />
              <Button
                variation={ButtonVariation.TERTIARY}
                text={getString('cancel')}
                size={ButtonSize.MEDIUM}
                onClick={() => setEdit(false)}
              />
            </Layout.Horizontal>
          </Container>
        </Truthy>
        <Else>
          <>
            <Text tag="h1" font={{ variation: FontVariation.H4 }}>
              {original} <span className={css.prNumber}>#{number}</span>
            </Text>
            <Button
              variation={ButtonVariation.ICON}
              tooltip={getString('edit')}
              tooltipProps={{ isDark: true, position: 'right' }}
              size={ButtonSize.SMALL}
              icon="code-edit"
              className={css.btn}
              onClick={() => setEdit(true)}
            />
            {!(description || '').trim().length && (
              <>
                <PipeSeparator height={10} />
                <a {...ButtonRoleProps} onClick={onAddDescriptionClick}>
                  &nbsp;{getString('pr.addDescription')}
                </a>
              </>
            )}
          </>
        </Else>
      </Match>
    </Layout.Horizontal>
  )
}