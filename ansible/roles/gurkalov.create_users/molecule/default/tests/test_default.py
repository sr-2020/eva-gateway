import os

import testinfra.utils.ansible_runner

testinfra_hosts = testinfra.utils.ansible_runner.AnsibleRunner(
    os.environ['MOLECULE_INVENTORY_FILE']).get_hosts('all')


def test_user1_file(host):
    f = host.file('/home/user1')

    assert f.exists


def test_user2_file(host):
    f = host.file('/home/user2')

    assert f.exists
