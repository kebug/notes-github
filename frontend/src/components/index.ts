import type { App } from 'vue'
import {
  create,
  NButton,
  NCard,
  NInput,
  NForm,
  NFormItem,
  NGrid,
  NGridItem,
  NSpace,
  NList,
  NListItem,
  NThing,
  NTag,
  NPagination,
  NPopconfirm,
  NIcon,
  NEmpty,
  NSpin,
  NDivider,
  NTooltip,
  NSwitch,
  NSelect,
  NDatePicker,
  NTimePicker,
  NInputNumber,
  NRadio,
  NRadioGroup,
  NCheckbox,
  NCheckboxGroup,
  NSlider,
  NRate,
  NUpload,
  NProgress,
  NLoadingBarProvider,
  NMessageProvider,
  NNotificationProvider,
  NDialogProvider
} from 'naive-ui'

const naive = create({
  components: [
    NButton,
    NCard,
    NInput,
    NForm,
    NFormItem,
    NGrid,
    NGridItem,
    NSpace,
    NList,
    NListItem,
    NThing,
    NTag,
    NPagination,
    NPopconfirm,
    NIcon,
    NEmpty,
    NSpin,
    NDivider,
    NTooltip,
    NSwitch,
    NSelect,
    NDatePicker,
    NTimePicker,
    NInputNumber,
    NRadio,
    NRadioGroup,
    NCheckbox,
    NCheckboxGroup,
    NSlider,
    NRate,
    NUpload,
    NProgress
  ]
})

export function setupNaive(app: App) {
  app.use(naive)
}

export {
  NLoadingBarProvider,
  NMessageProvider,
  NNotificationProvider,
  NDialogProvider
} 