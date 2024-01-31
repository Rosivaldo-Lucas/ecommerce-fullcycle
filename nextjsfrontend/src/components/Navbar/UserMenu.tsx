'use client';

import { useRouter } from "next/navigation";
import { useState } from "react";

import { IconButton, Menu, MenuItem, Typography } from '@mui/material';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import ShoppingCartIcon from '@mui/icons-material/ShoppingCart';
import LogoutIcon from '@mui/icons-material/Logout';
import ListAltIcon from '@mui/icons-material/ListAlt';

import Link from 'next/link';


export type UserMenuProps = {
  user: any | null;
};

export function UserMenu(props: UserMenuProps) {
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const router = useRouter();

  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const redirectToCart = () => {
    handleClose();
    router.push('/my-cart');
  };

  const redirectToMyOrders = () => {
    handleClose();
    router.push('/my-orders');
  };

  const handleLogout = async () => {
    handleClose();
  };

  return props.user ? (
    <div>
      <IconButton size='large' onClick={handleMenu}>
        <AccountCircleIcon />
      </IconButton>

      <Menu
        anchorEl={anchorEl}
        anchorOrigin={{
          vertical: 'top',
          horizontal: 'right'
        }}
        keepMounted
        transformOrigin={{
          vertical: 'top',
          horizontal: 'right'
        }}
        open={Boolean(anchorEl)}
        onClose={handleClose}
      >
        <MenuItem onClick={redirectToCart}>
          <ShoppingCartIcon />
          <Typography>Meu carrinho</Typography>
        </MenuItem>

        <MenuItem onClick={redirectToMyOrders}>
          <ListAltIcon />
          <Typography>Meus pedidos</Typography>
        </MenuItem>

        <MenuItem onClick={handleLogout}>
          <ListAltIcon />
          <Typography>Sair</Typography>
        </MenuItem>
      </Menu>
    </div>
  ) : (
    <Link href={'/login'} style={{ textDecoration: 'none' }}>
      <Typography color='text.primary' sx={{ ml: 3, fontWeight: 500 }}>
        Entrar
      </Typography>
    </Link>
  );
}