import { useStore } from '/@/store/index';

export function getI18nSource(key: string) {
  const store = useStore();
  const SysI18nSource = store.state.sysI18nSource.sysI18nSource
  if (!SysI18nSource || SysI18nSource[key] == undefined) {
    return key
  } else {
    return SysI18nSource[key]
  }
}