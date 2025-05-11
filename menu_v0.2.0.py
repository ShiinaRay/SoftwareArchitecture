import sys
import click

@click.group(context_settings={'help_option_names': ['-h', '--help']})
def cli():
    """通用命令行菜单子系统"""
    pass  # 顶层命令组:contentReference[oaicite:5]{index=5}

@cli.command()
def option_a():
    """执行选项 A"""
    click.echo("你选择了 A")  # 简洁输出:contentReference[oaicite:6]{index=6}

@cli.command()
@click.option('--count', '-c', default=1, help='重复次数')
def option_b(count):
    """执行选项 B（支持参数）"""
    for i in range(count):
        click.echo(f"B 执行第 {i+1} 次")  # 参数支持:contentReference[oaicite:7]{index=7}

@cli.command()
def version():
    """显示版本信息"""
    click.echo("menu v0.2.0")  # 子命令风格:contentReference[oaicite:8]{index=8}

def main():
    cli()

if __name__ == "__main__":
    main()  # 入口调用:contentReference[oaicite:9]{index=9}
