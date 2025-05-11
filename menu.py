import sys

def menu():
    print("=== 主菜单 ===")
    print("1. 选项 A")
    print("2. 选项 B")
    print("0. 退出")
    return input("请选择: ")

def action_a():
    print("你选择了 A")

def action_b():
    print("你选择了 B")

def main():
    while True:
        choice = menu()
        if choice == '1':
            action_a()
        elif choice == '2':
            action_b()
        elif choice == '0':
            print("再见！")
            sys.exit(0)
        else:
            print("无效选项，请重试。")

if __name__ == "__main__":
    main()
