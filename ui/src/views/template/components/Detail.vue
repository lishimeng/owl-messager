<template>
  <div class="createPost-container">
    <el-form ref="postForm" :model="postForm" :rules="rules" class="form-container">

      <sticky :z-index="10" :class-name="'sub-navbar '+postForm.status">
        <el-button v-loading="loading" style="margin-left: 10px;" type="success" @click="submitForm">
          Publish
        </el-button>
      </sticky>

      <div class="createPost-main-container">
        <el-row>
          <Warning />

          <el-col :span="24">
            <el-form-item style="margin-bottom: 40px;" prop="name">
              <MDinput v-model="postForm.name" :maxlength="100" name="name" required>
                Name
              </MDinput>
            </el-form-item>

            <div class="postInfo-container">
              <el-row>
                <el-col :span="8">
                  <el-form-item label-width="60px" label="类别:" class="postInfo-container-item">
                    <el-select v-model="postForm.category" filterable default-first-option remote placeholder="Select Category">
                      <el-option v-for="item in cateList" :key="item.category" :label="item.cateDesc" :value="item.category" />
                    </el-select>
                  </el-form-item>
                </el-col>
              </el-row>
            </div>
          </el-col>
        </el-row>

        <el-form-item style="margin-bottom: 40px;" label-width="70px" label="描述:">
          <el-input v-model="postForm.description" :rows="1" type="textarea" class="article-textarea" autosize placeholder="Please enter the description" />
          <span v-show="descriptionLength" class="word-counter">{{ descriptionLength }}words</span>
        </el-form-item>

        <el-form-item prop="body" style="margin-bottom: 30px;">
          <Tinymce ref="editor" v-model="postForm.body" :height="400" />
        </el-form-item>

      </div>
    </el-form>
  </div>
</template>

<script>
import Tinymce from '@/components/Tinymce'
import MDinput from '@/components/MDinput'
import Sticky from '@/components/Sticky' // 粘性header组件
import Warning from './Warning'
import { createMailTemplateApi } from '@/api/mail'

const defaultForm = {
  status: 'submit',
  name: '', // 文章题目
  body: '', // 文章内容
  description: '', // 文章摘要
  id: undefined,
  category: 0,
  platforms: ['a-platform'],
  comment_disabled: false
}

export default {
  name: 'Detail',
  components: { Tinymce, MDinput, Sticky, Warning },
  props: {
    isEdit: {
      type: Boolean,
      default: false
    }
  },
  data() {
    const validateRequire = (rule, value, callback) => {
      if (value === '') {
        this.$message({
          message: rule.field + '为必传项',
          type: 'error'
        })
        callback(new Error(rule.field + '为必传项'))
      } else {
        callback()
      }
    }
    return {
      postForm: Object.assign({}, defaultForm),
      loading: false,
      rules: {
        name: [{ validator: validateRequire }],
        body: [{ validator: validateRequire }]
      },
      tempRoute: {},
      cateList: [
        { category: 0, cateDesc: '默认' },
        { category: 1, cateDesc: '短信' },
        { category: 2, cateDesc: '邮件' },
        { category: 3, cateDesc: '其它' }
      ]
    }
  },
  computed: {
    descriptionLength() {
      return this.postForm.description.length
    },
    displayTime: {
      // set and get is useful when the data
      // returned by the back end api is different from the front end
      // back end return => "2013-06-25 06:59:25"
      // front end need timestamp => 1372114765000
      get() {
        return (+new Date(this.postForm.display_time))
      },
      set(val) {
        this.postForm.display_time = new Date(val)
      }
    }
  },
  created() {
    // if (this.isEdit) {
    //   const id = this.$route.params && this.$route.params.id
    //   this.fetchData(id)
    // }

    // // Why need to make a copy of this.$route here?
    // // Because if you enter this page and quickly switch tag, may be in the execution of the setTagsViewTitle function, this.$route is no longer pointing to the current page
    // // https://github.com/PanJiaChen/vue-element-admin/issues/1221
    // this.tempRoute = Object.assign({}, this.$route)
  },
  methods: {
    createMailTemplate() {
      createMailTemplateApi({
        name: this.postForm.name,
        body: this.postForm.body,
        description: this.postForm.description,
        category: this.postForm.category
      }).then(res => {
        if (res && res.code !== -1) {
          this.$notify({
            title: '成功',
            message: '发布文章成功',
            type: 'success',
            duration: 2000
          })
          this.postForm.status = 'published'
          this.loading = false
        }
      })
    },
    submitForm() {
      console.log(this.postForm)
      this.$refs.postForm.validate(valid => {
        if (valid) {
          this.loading = true
          this.createMailTemplate()
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
@import "~@/styles/mixin.scss";

.createPost-container {
  position: relative;

  .createPost-main-container {
    padding: 40px 45px 20px 50px;

    .postInfo-container {
      position: relative;
      @include clearfix;
      margin-bottom: 10px;

      .postInfo-container-item {
        float: left;
      }
    }
  }

  .word-counter {
    width: 40px;
    position: absolute;
    right: 10px;
    top: 0px;
  }
}

.article-textarea ::v-deep {
  textarea {
    padding-right: 40px;
    resize: none;
    border: none;
    border-radius: 0px;
    border-bottom: 1px solid #bfcbd9;
  }
}
</style>
