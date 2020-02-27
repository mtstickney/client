import * as React from 'react'
import {SectionList, SectionListRenderItem as _SectionListRenderItem} from 'react-native'

// Desktop specific props. `selectedIndex` is used for SectionList with item
// selecting, where the scroll should follow selected item.
type DesktopProps = {
  selectedIndex?: number
  disableAbsoluteStickyHeader?: boolean
}

export type Props<T> = React.ComponentProps<SectionList<T>> & DesktopProps
export type SectionListRenderItem<T> = _SectionListRenderItem<T>

export default class<T> extends React.Component<Props<T>> {}
