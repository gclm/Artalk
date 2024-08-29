import type { MessageSchema } from '../i18n'

export const zhTW: MessageSchema = {
  ctrlCenter: '控制中心',
  msgCenter: '訊息中心',
  noContent: '無內容',
  searchHint: '關鍵字搜尋...',
  allSites: '所有網站',
  siteManage: '網站管理',
  comment: '評論',
  page: '頁面',
  user: '用戶',
  site: '網站',
  transfer: '轉移',
  settings: '設置',
  all: '全部',
  pending: '待審',
  personal: '個人',
  mentions: '提及',
  mine: '我的',
  admin: '管理員',
  create: '建立',
  import: '匯入',
  export: '匯出',
  settingSaved: '設定已儲存',
  settingSaveFailed: '設定儲存失敗',
  settingNotice: '注：某些設定選項可能需要手動重啟才能生效',
  apply: '套用',
  updateComplete: '更新完畢',
  updateReady: '準備更新...',
  opFailed: '操作失敗',
  updateTitle: '擷取標題',
  uploading: '上傳中',
  cancel: '取消',
  back: '返回',
  cacheClear: '清除快取',
  cacheWarm: '預熱快取',
  editTitle: '編輯標題',
  switchKey: 'KEY 變更',
  commentAllowAll: '允許任何人評論',
  commentOnlyAdmin: '僅允許管理員評論',
  config: '配置文件',
  envVarControlHint: '由環境變數 {key} 參照',
  userAdminHint: '該用戶具有管理員權限',
  userInConfHint: '該用戶存在於配置文件中',
  userInConfCannotEditHint: '暫不支持線上編輯配置文件中的用戶，請手動修改配置文件',
  userDeleteConfirm:
    '該操作將刪除 用戶："{name}" 郵箱："{email}" 所有評論，包括其評論下面他人的回覆評論，是否繼續？',
  userDeleteManuallyHint: '用戶已從數據庫刪除，請手動編輯配置文件並刪除用戶',
  pageDeleteConfirm: '確認刪除頁面 "{title}"？將會刪除所有相關數據',
  siteDeleteConfirm: '該操作將刪除網站："{name}" 及其下所有數據，是否繼續？',
  siteNameInputHint: '請輸入網站名稱',
  edit: '編輯',
  delete: '刪除',
  siteCount: '共 {count} 個網站',
  createSite: '建立網站',
  siteName: '網站名稱',
  siteUrls: '網站 URLs',
  multiSepHint: '使用逗號分隔多個',
  add: '新增',
  rename: '重命名',
  inputHint: '輸入文字...',
  userCreate: '建立用戶',
  userEdit: '用戶編輯',
  comments: '評論',
  last: '最後',
  show: '展開',
  username: '用戶名',
  email: '郵箱',
  link: '連結',
  badgeText: '徽章文字',
  badgeColor: '徽章顏色',
  role: '身份角色',
  normal: '一般',
  password: '密碼',
  passwordEmptyHint: '留空表示不變更密碼',
  emailNotify: '郵件通知',
  enabled: '啟用',
  disabled: '停用',
  save: '儲存',
  dataFile: '資料檔案',
  artransfer: '轉移工具',
  targetSiteName: '目標網站名稱',
  targetSiteURL: '目標網站 URL',
  payload: '有效載荷',
  optional: '選填',
  uploadReadyToImport: '檔案已上傳並準備匯入',
  artransferToolHint: '使用 {link} 將評論數據轉換為 Artrans 格式。',
  moreDetails: '查看詳情',
  loginFailure: '登入失敗',
  login: '登入',
  logout: '登出',
  logoutConfirm: '確定要登出嗎？',
  loginSelectHint: '請選擇您要登入的帳號：',
}

export default zhTW