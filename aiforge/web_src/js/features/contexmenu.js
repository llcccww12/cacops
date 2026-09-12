
export default async function initContextMenu() {
    $('.popup-close').on('click', function (e) {
        $(this).parents('tr').prev().css('display', 'table-row')
        $(this).closest('tr').css('cssText', 'display:none !important')
    })
    function contextMenu() {
        let canContextMenu = $('.ui.single.line.table.can-context-menu').data('can-editfile')
        if (canContextMenu) {
            $('.name.four.wide').on('contextmenu', function (e) {
                let ev = window.event || e;
                ev.preventDefault();
                menu.show(e)
            })
        } else {
            return
        }
    }
    contextMenu()
    const menu = new Menu({
        data: [
            {
                label: '新标签打开',
                icon: "file outline icon context-menu-icon",
                active: (e, a) => {
                    window.open(a.currentTarget.getElementsByTagName("a")[0].getAttribute('href'))
                }
            },
            {
                label: '重命名',
                icon: "edit icon context-menu-icon",
                active: (e, a) => {
                    document.querySelectorAll(".context-menu-one").forEach((ele) => {
                        if (ele.style.display === 'table-row') {

                            ele.style.display = 'none'
                            ele.previousElementSibling.style.display = 'table-row'
                        }
                    })
                    if (a.currentTarget.parentNode.nextElementSibling) {
                        a.currentTarget.parentNode.style.setProperty('display', 'none', 'important')
                        a.currentTarget.parentNode.nextElementSibling.style.display = 'table-row'
                        const renameFile = a.currentTarget.getElementsByTagName("a")[0].getAttribute('title')
                        let renameFileValue = renameFile.indexOf('/') !== -1 ? renameFile.substr(0, renameFile.indexOf('/')) : renameFile
                        a.currentTarget.parentNode.nextElementSibling.getElementsByTagName("input")[0].setAttribute("value", renameFileValue)

                    }
                    let btn = a.currentTarget.parentNode.nextElementSibling.getElementsByTagName("button")[0]
                    btn.addEventListener('click', function (e) {
                        let postUrl = btn.getAttribute('data-postbasepath')
                        const postUrlArr = postUrl.split('/')
                        postUrlArr[postUrlArr.length - 1] = encodeURIComponent(postUrlArr[postUrlArr.length - 1])
                        postUrl = postUrlArr.join('/')
                        console.log(postUrl)
                        let last_commit = btn.getAttribute('data-commit')
                        let tree_path = btn.getAttribute('data-treepath') + e.target.parentNode.previousElementSibling.getElementsByTagName("input")[0].value
                        let csrf = $("input[name='_csrf']").val()
                        $.ajax({
                            url: postUrl,
                            type: "POST",
                            contentType: "application/x-www-form-urlencoded",
                            data: { last_commit: last_commit, tree_path: tree_path, _csrf: csrf },
                            success: function (res) {
                                if (res.Code === 0) {
                                    document.getElementById("mask").style.display = "block"
                                    location.reload()
                                }
                                else {
                                    $('.ui.negative.message').text(res.Msg).show().delay(10000).fadeOut();
                                }
                            }

                        })

                    })
                },
            },
            // {
            // 	label: '删除',
            // 	icon: "trash icon context-menu-icon",
            // 	active: (e, a) => {
            // 		console.dir(a)
            // 		$('.context-menu-delete.modal')
            // 			.modal({
            // 				onApprove() {
            // 				}
            // 			})
            // 			.modal('show');
            // 	}
            // },

        ]
    })
}
class Menu {
    constructor(param) {
        this.target = document.createElement('div')
        this.target.classList.add("ui", "menu", "compact", "vertical", "context-menu-click")
        this.data = param.data
        this.active = false
        this.clickZ = this.click.bind(this)
        this.closeZ = this.close2.bind(this)
        document.addEventListener('click', this.closeZ)
        for (let i = 0; i < this.data.length; i++) {
            let div = document.createElement('a')
            div.classList.add('item', 'context-menu-operation')
            if (this.data[i].disabled) {
                div.classList.add('disabled')
            }
            div.dataset.index = i.toString()
            div.innerHTML = `<i class="${this.data[i].icon}"></i>${this.data[i].label}`
            div.addEventListener('click', this.clickZ)
            this.target.append(div)
        }
        document.body.append(this.target)
    }
    click(e) {
        let index = parseInt(e.target.dataset.index)
        if (this.data[index].active) {
            this.data[index].active(e, this.acEvent)
        }
        this.close()
    }
    show(e) {
        this.active = true
        this.nodeList = this.target.querySelectorAll('.item')
        for (let i = 0; i < this.data.length; i++) {
            if (this.data[i].beforeDisabled) {
                let t = this.data[i].beforeDisabled(e)
                this.data[i].disabled = t
                if (t) {
                    this.nodeList[i].classList.add('disabled')
                } else {
                    this.nodeList[i].classList.remove('disabled')
                }
            }
        }
        this.acEvent = e
        this.target.style.top = `${e.pageY}px`
        this.target.style.left = `${e.pageX}px`
        this.target.style.minWidth = '90px'
        this.target.classList.add('active')
    }

    close() {
        this.active = false
        this.target.classList.remove('active')
    }

    close2(e) {
        if (!this.target.contains(e.target) && this.active) {
            this.active = false
            this.target.classList.remove('active')
        }
    }
}
