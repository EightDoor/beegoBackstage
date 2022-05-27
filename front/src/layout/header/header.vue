<template>
  <a-layout-header class="container">
    <div class="content">
      <div class="button-icon-container">
        <div @click="ToggleCollapsed()">
          <MenuUnfoldOutlined v-if="collapsed" class="button-icon" />
          <MenuFoldOutlined v-else class="button-icon" />
        </div>
        <a-breadcrumb>
          <a-breadcrumb-item
            v-for="(item, index) in crumbs"
            :key="index"
            class="button-breadcrumb"
          >
            {{ item }} {{ index > 0 && index + 1 !== crumbs.length ? '/' : '' }}
          </a-breadcrumb-item>
        </a-breadcrumb>
      </div>
      <a-dropdown>
        <a class="ant-dropdown-link" @click="(e) => e.preventDefault()">
          <span v-if="userInfo" class="nickName">
            {{ userInfo.nickName }}
          </span>
          <template v-if="userInfo.file">
            <a-avatar :src="userInfo.file.url" :size="50" />
          </template>
          <a-avatar v-else :size="50">
            <template #icon><UserOutlined /></template>
          </a-avatar>
        </a>
        <template #overlay>
          <a-menu>
            <a-menu-item v-for="(item, index) in data" :key="index">
              <a @click="GoTo(item)">{{ item }}</a>
            </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
    </div>
  </a-layout-header>
  <CommonDrawer :footer="false" :visible="visible" title="个人中心" @on-close="visible = false">
    <a-form>
      <a-form-item label="旧密码" placeholder="请输入旧密码" v-bind="validateInfos.password">
        <a-input v-model:value="modelRef.password" />
      </a-form-item>
      <a-form-item label="新密码" placeholder="请输入新密码" v-bind="validateInfos.newPassword">
        <a-input v-model:value="modelRef.newPassword" />
      </a-form-item>
      <a-form-item>
        <a-button type="primary" @click.prevent="onSubmit()">
          更改
        </a-button>
      </a-form-item>
    </a-form>
  </CommonDrawer>
</template>

<script lang="ts">
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  UserOutlined,
} from '@ant-design/icons-vue'
import { computed, defineComponent, reactive, ref, toRaw } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { Form, message } from 'ant-design-vue'
import { COLLAPSED } from '@/store/mutation-types'
import log from '@/utils/log'
import CommonDrawer from '@/components/Drawer/Drawer.vue'
import storeInstant from '@/utils/store'
import http from '@/utils/request'
import type { UserType } from '@/types/sys'

export default defineComponent({
  name: 'CommonHeader',
  components: {
    UserOutlined,
    MenuUnfoldOutlined,
    MenuFoldOutlined,
    CommonDrawer,
  },
  setup() {
    const data = ref<string[]>(['个人中心', '退出'])
    const router = useRouter()
    const store = useStore()
    const userInfo = computed<UserType>(() => store.state.sys.userInfo)

    // 个人中心
    const visible = ref(false)
    const { useForm } = Form
    const modelRef = reactive({
      password: '',
      newPassword: '',
      id: 0,
    })
    const ruleRef = ref({
      password: [
        {
          required: true,
          message: '请输入旧密码',
        },
      ],
      newPassword: [
        {
          required: true,
          message: '请输入新密码',
        },
      ],
    })

    const { resetFields, validate, validateInfos } = useForm(modelRef, ruleRef)

    function onSubmit() {
      validate()
        .then(() => {
          const data = toRaw(modelRef)
          data.id = userInfo.value.id!
          log.i(data, '表单值')
          http({
            url: 'user/password',
            method: 'POST',
            body: data,
          }).then((res) => {
            const data = res.data
            log.i(data, '修改返回的值')
            message.success('修改成功')
            resetFields()
          })
        })
        .catch((err) => {
          log.e(err, '表单验证错误')
        })
    }

    function GoTo(val: string) {
      switch (val) {
        case '退出':
          localStorage.clear()
          storeInstant.clear()
          router.replace('/login')
          break
        case '个人中心':
          visible.value = true
          break
        default:
      }
    }
    function ToggleCollapsed() {
      store.commit(COLLAPSED)
    }
    const crumbs = computed(() => store.state.crumbs.list)

    return {
      // data
      data,
      collapsed: computed(() => store.state.sys.collapsed),
      crumbs,
      // methods
      GoTo,
      ToggleCollapsed,

      visible,
      validateInfos,
      resetFields,
      modelRef,
      onSubmit,
      userInfo,
    }
  },
})
</script>

<style scoped lang="less">
@import 'header.less';
</style>
