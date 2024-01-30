import {
  Controller,
  Get,
  Post,
  Body,
  Param,
  UseGuards,
  Req,
} from '@nestjs/common';
import { OrdersService } from './orders.service';
import { CreateOrderDto } from './dto/create-order.dto';
import { AuthGuard } from 'src/auth/auth.guard';
import { Request } from 'express';

@UseGuards(AuthGuard)
@Controller('orders')
export class OrdersController {
  constructor(private readonly ordersService: OrdersService) {}

  @Post()
  create(@Body() createOrderDto: CreateOrderDto, @Req() request: Request) {
    return this.ordersService.create({
      ...createOrderDto,
      client_id: request['user'].sub,
    });
  }

  @Get()
  findAll(@Req() request: Request) {
    return this.ordersService.findAll(request['user'].sub);
  }

  @Get(':id')
  findOne(@Param('id') id: string, @Req() request: Request) {
    return this.ordersService.findOne(id, request['user'].sub);
  }
}
