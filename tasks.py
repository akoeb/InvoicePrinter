"""
This uses [http://docs.pyinvoke.org/](invoke) as command runner
"""
from invoke import task

args = ('-ti --rm --user 1000:1000 '
        ' -v "${PWD}":/opt/app '
        ' -v ${HOME}/.gitconfig:/.gitconfig '
        ' -w /opt/app '
        ' -e GOPATH=/opt/app '
        ' -p 8080:8080')

img = 'golang:1.6'


@task
def build(ctx):
    ctx.run('docker run {args} {img} go build -v'.format(args=args, img=img))


@task
def bash(ctx):
    ctx.run('docker run {args} {img} bash'.format(args=args, img=img), pty=True)


@task
def run(ctx, file=None):
    if file is None:
        ctx.run('docker run {args} {img} go run *.go'.format(args=args, img=img))
    else:
        ctx.run('docker run {args} {img} go run {file}'.format(args=args, img=img, file=file))


@task
def get(ctx, pkg=''):
    ctx.run('docker run {args} {img} go get -u -v {pkg}'.format(args=args, img=img, pkg=pkg))


@task
def test(ctx):
    ctx.run('docker run {args} {img} go test -v '.format(args=args, img=img))

@task
def pdf(ctx):
    ctx.run('docker run {args} {img} go run pdfwriter.go '.format(args=args, img=img))
