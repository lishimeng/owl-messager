import { useStore } from '/@/store/index';

export function getJsonConfigs(key: string) {
  const store = useStore();
  const jsonConfigs = store.state.jsonConfigs.jsonConfigs
  if (!jsonConfigs || jsonConfigs[key] == undefined) {
    return true
  } else {
    return jsonConfigs[key]
  }
}