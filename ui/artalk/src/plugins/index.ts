import type { ArtalkPlugin } from '@/types'
import { ConfRemoter } from './conf-remoter'
import { Markdown } from './markdown'
import { EditorKit } from './editor-kit'
import { ListPlugins } from './list'
import { Notifies } from './notifies'
import { PvCountWidget } from './stat'
import { VersionCheck } from './version-check'
import { AdminOnlyElem } from './admin-only-elem'
import { DarkMode } from './dark-mode'

export const DefaultPlugins: ArtalkPlugin[] = [
  ConfRemoter,
  Markdown, EditorKit, AdminOnlyElem,
  ...ListPlugins,
  Notifies,
  PvCountWidget, VersionCheck, DarkMode,
]
